//go build -ldflags "-s -w" MonitorServer.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

const MonitorServerRemark = `MonitorServer execute shell:
	MonitorServer [port=2019]

Example:
	MonitorServer 2019`

func getPort() string {
	if len(os.Args) > 1 {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(MonitorServerRemark)
		}
		return os.Args[1]
	}
	return "2019"
}

func loggingHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		str := fmt.Sprintf(
			"%s Comleted %s %s in %v from %s\n",
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			time.Since(start),
			r.RemoteAddr)
		go print(str)
	}
}

const (
	Byte  = 1
	KByte = Byte * 1024
	MByte = KByte * 1024
	GByte = MByte * 1024
	TByte = GByte * 1024
	PByte = TByte * 1024
	EByte = PByte * 1024
)

var bytesSizeTable = map[string]uint64{
	"b":  Byte,
	"kb": KByte,
	"mb": MByte,
	"gb": GByte,
	"tb": TByte,
	"pb": PByte,
	"eb": EByte,
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize calculates the file size and generate user-friendly string.
func FileSize(s int64) string {
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(uint64(s), 1024, sizes)
}

var startTime = time.Now()

type SysStatus struct {
	Uptime       string
	NumGoroutine int

	// General statistics.
	MemAllocated string // bytes allocated and still in use
	MemTotal     string // bytes allocated (even if freed)
	MemSys       string // bytes obtained from system (sum of XxxSys below)
	Lookups      uint64 // number of pointer lookups
	MemMallocs   uint64 // number of mallocs
	MemFrees     uint64 // number of frees

	// Main allocation heap statistics.
	HeapAlloc    string // bytes allocated and still in use
	HeapSys      string // bytes obtained from system
	HeapIdle     string // bytes in idle spans
	HeapInuse    string // bytes in non-idle span
	HeapReleased string // bytes released to the OS
	HeapObjects  uint64 // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string // bootstrap stacks
	StackSys    string
	MSpanInuse  string // mspan structures
	MSpanSys    string
	MCacheInuse string // mcache structures
	MCacheSys   string
	BuckHashSys string // profiling bucket hash table
	GCSys       string // GC metadata
	OtherSys    string // other system allocations

	// Garbage collector statistics.
	NextGC       string // next run in HeapAlloc time (bytes)
	LastGC       string // last run in absolute time (ns)
	PauseTotalNs string
	PauseNs      string // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
	NumGC        uint32
}

func NewSysStatus() *SysStatus {
	sysStatus := new(SysStatus)

	sysStatus.Uptime = time.Since(startTime).String()

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.NumGoroutine = runtime.NumGoroutine()

	sysStatus.MemAllocated = FileSize(int64(m.Alloc))
	sysStatus.MemTotal = FileSize(int64(m.TotalAlloc))
	sysStatus.MemSys = FileSize(int64(m.Sys))
	sysStatus.Lookups = m.Lookups
	sysStatus.MemMallocs = m.Mallocs
	sysStatus.MemFrees = m.Frees

	sysStatus.HeapAlloc = FileSize(int64(m.HeapAlloc))
	sysStatus.HeapSys = FileSize(int64(m.HeapSys))
	sysStatus.HeapIdle = FileSize(int64(m.HeapIdle))
	sysStatus.HeapInuse = FileSize(int64(m.HeapInuse))
	sysStatus.HeapReleased = FileSize(int64(m.HeapReleased))
	sysStatus.HeapObjects = m.HeapObjects

	sysStatus.StackInuse = FileSize(int64(m.StackInuse))
	sysStatus.StackSys = FileSize(int64(m.StackSys))
	sysStatus.MSpanInuse = FileSize(int64(m.MSpanInuse))
	sysStatus.MSpanSys = FileSize(int64(m.MSpanSys))
	sysStatus.MCacheInuse = FileSize(int64(m.MCacheInuse))
	sysStatus.MCacheSys = FileSize(int64(m.MCacheSys))
	sysStatus.BuckHashSys = FileSize(int64(m.BuckHashSys))
	sysStatus.GCSys = FileSize(int64(m.GCSys))
	sysStatus.OtherSys = FileSize(int64(m.OtherSys))

	sysStatus.NextGC = FileSize(int64(m.NextGC))
	sysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.NumGC = m.NumGC
	return sysStatus
}

func TextMarshal(s *SysStatus) string {
	data := make(map[string]interface{})
	v_elem := reflect.ValueOf(s).Elem()
	t_elem := reflect.TypeOf(s).Elem()
	for i := 0; i < v_elem.NumField(); i++ {
		v := v_elem.Field(i)
		data[t_elem.Field(i).Name] = v.Interface()
	}
	return fmt.Sprint(data)
}

func MonitorHandler(w http.ResponseWriter, r *http.Request) {
	var data []byte
	var err error
	switch r.URL.Path {
	case "/monitor.json":
		data, err = json.Marshal(NewSysStatus())
	case "/monitor.xml":
		data, err = xml.Marshal(NewSysStatus())
	case "/monitor.txt":
		data = []byte(TextMarshal(NewSysStatus()))
	default:
		err = errors.New("Not Found")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
	} else {
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(r.URL.Path)))
		w.Write(data)
	}
}

func main() {
	port := getPort()
	fmt.Println("> Listening at http://127.0.0.1:" + port)
	http.HandleFunc("/", loggingHandler(MonitorHandler))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		println(err.Error())
	}
}
