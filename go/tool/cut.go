//go build -ldflags "-s -w" cut.go
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const usage = `cut a slice from the head or tail of a (big) file, in parallel
cut (tail|head) N(k|m|g) file

Output: same dir, named <base>-<size>-<timestamp><ext>
	e.g. app.log + tail 1g -> app-1g-20261119153000.log

Units: k=1024, m=1024k, g=1024m (decimals allowed, e.g. 1.5g)
If N is larger than the file, the whole file is copied.

Performance & memory:
	The range is split into chunks copied by concurrent goroutines
	via ReadAt/WriteAt. Peak memory is bounded and constant:
	workers x chunk (default 4 x 4MB = 16MB), no matter if the
	file is 1MB or 100GB. tail jumps straight to the cut point.

Tuning (env):
	CUT_WORKERS=8        parallel goroutines (default min(NumCPU,4);
	                     use 1 for spinning disks)
	CUT_CHUNK=8m         chunk size per worker (default 4m, max 256m)

Example:
	cut tail 1g app.log
	cut head 500m /var/log/big.log
`

// parseSize parses "100k", "1.5g", "512m" (or plain bytes) into a byte count.
func parseSize(s string) (int64, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return 0, fmt.Errorf("empty size")
	}
	mult := int64(1)
	num := s
	switch s[len(s)-1] {
	case 'k':
		mult, num = 1024, s[:len(s)-1]
	case 'm':
		mult, num = 1024*1024, s[:len(s)-1]
	case 'g':
		mult, num = 1024*1024*1024, s[:len(s)-1]
	}
	v, err := strconv.ParseFloat(num, 64)
	if err != nil || v <= 0 {
		return 0, fmt.Errorf("invalid size: %s", s)
	}
	n := int64(v * float64(mult))
	if n <= 0 {
		return 0, fmt.Errorf("size too small: %s", s)
	}
	return n, nil
}

// config returns the worker count and chunk size (env-overridable).
// Peak copy memory is workers*chunkSize, independent of the file size.
func config() (int, int64) {
	workers := runtime.NumCPU()
	if workers > 4 {
		workers = 4
	}
	if v := os.Getenv("CUT_WORKERS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			workers = n
		}
	}
	if workers > 64 {
		workers = 64
	}
	chunk := int64(4 << 20)
	if v := os.Getenv("CUT_CHUNK"); v != "" {
		if n, err := parseSize(v); err == nil && n > 0 {
			chunk = n
		}
	}
	if chunk > 256<<20 {
		chunk = 256 << 20
	}
	return workers, chunk
}

// createOut creates <base>-<size>-<timestamp><ext> next to the source file.
// O_EXCL guarantees we never overwrite an existing file; on a name clash
// (same second) a -2, -3 ... suffix is appended before the extension.
func createOut(src, sizeArg string) (string, *os.File, error) {
	ext := filepath.Ext(src)
	base := strings.TrimSuffix(filepath.Base(src), ext)
	stem := base + "-" + strings.ToLower(strings.TrimSpace(sizeArg)) + "-" + time.Now().Format("20060102150405")
	for i := 0; ; i++ {
		name := stem + ext
		if i > 0 {
			name = stem + "-" + strconv.Itoa(i+1) + ext
		}
		path := filepath.Join(filepath.Dir(src), name)
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err == nil {
			return path, f, nil
		}
		if !os.IsExist(err) {
			return "", nil, err
		}
	}
}

func human(n int64) string {
	switch {
	case n >= 1<<30:
		return strconv.FormatFloat(float64(n)/(1<<30), 'f', 2, 64) + " GB"
	case n >= 1<<20:
		return strconv.FormatFloat(float64(n)/(1<<20), 'f', 2, 64) + " MB"
	case n >= 1<<10:
		return strconv.FormatFloat(float64(n)/(1<<10), 'f', 2, 64) + " KB"
	default:
		return strconv.FormatInt(n, 10) + " B"
	}
}

type chunk struct {
	srcOff, dstOff int64
	size           int
}

// transfer copies one chunk: ReadAt/WriteAt are offset-based, so any number
// of goroutines can hit the same files concurrently without sharing a seek
// position. The buffer is owned by the worker and reused for every chunk.
func transfer(in, out *os.File, buf []byte, c chunk) error {
	off := 0
	for off < c.size {
		nr, err := in.ReadAt(buf[off:c.size], c.srcOff+int64(off))
		off += nr
		if off == c.size {
			break
		}
		if err != nil {
			return fmt.Errorf("read at %d: %w", c.srcOff+int64(off), err)
		}
		if nr == 0 {
			return io.ErrNoProgress
		}
	}
	off = 0
	for off < c.size {
		nw, err := out.WriteAt(buf[off:c.size], c.dstOff+int64(off))
		off += nw
		if err != nil {
			return fmt.Errorf("write at %d: %w", c.dstOff+int64(off), err)
		}
		if nw == 0 {
			return io.ErrShortWrite
		}
	}
	return nil
}

// copyParallel streams n bytes starting at offset from in to out using a
// fixed pool of workers. Memory is bounded by workers*chunkSize; the file
// itself is never loaded. Chunks land at their exact output offsets, so no
// ordering or reassembly is needed.
func copyParallel(in, out *os.File, offset, n int64, workers int, chunkSize int64) (int64, error) {
	// Pre-size the output (instant on NTFS/ext4, no zero-fill). Workers
	// then WriteAt inside the final extent instead of racing to extend it.
	if err := out.Truncate(n); err != nil {
		return 0, err
	}

	var done int64
	var printed int32
	var mu sync.Mutex
	var firstErr error
	var once sync.Once
	quit := make(chan struct{})
	fail := func(err error) {
		mu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		mu.Unlock()
		once.Do(func() { close(quit) })
	}

	jobs := make(chan chunk, workers)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			buf := make([]byte, chunkSize) // one fixed buffer per worker
			for {
				select {
				case <-quit:
					return
				case c, ok := <-jobs:
					if !ok {
						return
					}
					if err := transfer(in, out, buf, c); err != nil {
						fail(err)
						return
					}
					atomic.AddInt64(&done, int64(c.size))
				}
			}
		}()
	}

	// Progress line twice a second; reads an atomic counter, allocates nothing.
	stopProg := make(chan struct{})
	progDone := make(chan struct{})
	go func() {
		defer close(progDone)
		tick := time.NewTicker(500 * time.Millisecond)
		defer tick.Stop()
		last := int64(-1)
		for {
			select {
			case <-stopProg:
				return
			case <-tick.C:
				d := atomic.LoadInt64(&done)
				if d != last && d > 0 {
					fmt.Fprintf(os.Stderr, "\r  %s / %s (%d%%)", human(d), human(n), d*100/n)
					atomic.StoreInt32(&printed, 1)
					last = d
				}
			}
		}
	}()

	total := (n + chunkSize - 1) / chunkSize
feed:
	for i := int64(0); i < total; i++ {
		size := chunkSize
		if rem := n - i*chunkSize; rem < size {
			size = rem
		}
		c := chunk{srcOff: offset + i*chunkSize, dstOff: i * chunkSize, size: int(size)}
		select {
		case jobs <- c:
		case <-quit:
			break feed
		}
	}
	close(jobs)
	wg.Wait()
	close(stopProg)
	<-progDone
	if atomic.LoadInt32(&printed) == 1 {
		fmt.Fprintln(os.Stderr)
	}
	if firstErr != nil {
		return atomic.LoadInt64(&done), firstErr
	}
	return n, nil
}

func run() error {
	if len(os.Args) != 4 {
		fmt.Print(usage)
		os.Exit(1)
	}
	mode := strings.ToLower(os.Args[1])
	if mode != "tail" && mode != "head" {
		return fmt.Errorf("first arg must be tail or head, got: %s", os.Args[1])
	}
	size, err := parseSize(os.Args[2])
	if err != nil {
		return err
	}
	src := os.Args[3]

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	stat, err := in.Stat()
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("not a file: %s", src)
	}

	n := size
	if n > stat.Size() {
		n = stat.Size()
	}
	offset := int64(0)
	if mode == "tail" {
		offset = stat.Size() - n // O(1) jump, the leading bytes are never read
	}

	dst, out, err := createOut(src, os.Args[2])
	if err != nil {
		return err
	}
	workers, chunkSize := config()
	written, err := copyParallel(in, out, offset, n, workers, chunkSize)
	if cerr := out.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		os.Remove(dst)
		return err
	}
	fmt.Printf("%s %s of %s (%d bytes, %d workers x %s chunks)\n-> %s\n",
		mode, os.Args[2], src, written, workers, human(chunkSize), dst)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "cut:", err)
		os.Exit(1)
	}
}
