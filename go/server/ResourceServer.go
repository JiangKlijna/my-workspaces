package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Parameter struct {
	Dir  string
	Port string
	File string
}

type StaticFile struct {
	isDir      bool
	name, path string
}

var paras = &Parameter{}

func init() {
	var help bool
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&(paras.Dir), "d", "./", "static resource folder")
	flag.StringVar(&(paras.Port), "p", "80", "listening port")
	flag.StringVar(&(paras.File), "o", "static_gen.go", "golang source file")
	flag.Parse()
	if help {
		println(`Usage: [-p port] [-d dir] [-o file]
Example: -p 80 -d . -o static_gen.go

Options:`)
		flag.PrintDefaults()
		os.Exit(1)
	}
	absdir, err := filepath.Abs(paras.Dir)
	if err != nil {
		log.Fatal(err)
	}
	paras.Dir = absdir
}

// generate gen_go_file
func gen(files chan *StaticFile) {
	buf := bytes.Buffer{}
	buf.WriteString(`package main
import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"
)
var modtime = time.Now()
type MemoryFile struct {
	*bytes.Reader
	size  int64
	name  string
	isDir bool
}
func (m MemoryFile) Close() error {
	return nil
}
func (m MemoryFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, errors.New("no dir")
}
func (m MemoryFile) Stat() (os.FileInfo, error) {
	return m, nil
}
func (m MemoryFile) Name() string {
	return m.name
}
func (m MemoryFile) Size() int64 {
	return m.size
}
func (m MemoryFile) Mode() os.FileMode {
	return os.ModePerm
}
func (m MemoryFile) ModTime() time.Time {
	return modtime
}
func (m MemoryFile) IsDir() bool {
	return m.isDir
}
func (m MemoryFile) Sys() interface{} {
	return nil
}
type FakeFileSystem struct {
}
func (ffs FakeFileSystem) Open(name string) (http.File, error) {
	if data, ok := R[name]; ok {
		if data != nil {
			return &MemoryFile{bytes.NewReader(data), int64(len(data)), name, false}, nil
		} else {
			return &MemoryFile{nil, 0, name, true}, nil
		}
	} else {
		return nil, syscall.ERROR_PATH_NOT_FOUND
	}
}
func main() {
	http.Handle("/", http.FileServer(&FakeFileSystem{}))
	err := http.ListenAndServe(":` + paras.Port + `", nil)
	log.Fatal(err)
}
var R = map[string][]byte{`)
	for {
		sf := <-files
		if sf == nil {
			break
		}
		if sf.isDir {
			buf.WriteString("\n\t\"/" + sf.path)
			buf.WriteString(`":nil,`)
		} else {
			bs, err := ioutil.ReadFile(sf.name)
			if err != nil {
				panic(err)
			}
			buf.WriteString("\n\t\"/" + sf.path)
			buf.WriteString(`":{`)
			for _, b := range bs {
				buf.WriteString(strconv.Itoa(int(b)))
				buf.WriteString(",")
			}
			buf.WriteString("},")
		}
		fmt.Println(sf.name + " generate successful")
	}
	buf.WriteString("\n}")
	err := ioutil.WriteFile(paras.File, buf.Bytes(), 0662)
	if err != nil {
		panic(err)
	}
	files <- nil
}

func getFiles(dir string, files chan *StaticFile) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fs {
		name := dir + "/" + f.Name()
		files <- &StaticFile{f.IsDir(), name, name[len(paras.Dir)+1:]}
		if f.IsDir() {
			getFiles(name, files)
		}
	}
}

func main() {
	files := make(chan *StaticFile)
	go gen(files)              // 开始生成
	getFiles(paras.Dir, files) // 遍历所有文件
	files <- nil               // 结束遍历
	<-files                    // 结束生成程序
}
