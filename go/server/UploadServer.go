//go build -ldflags "-s -w" UploadServer.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

const UploadServerRemark = `UploadServer execute shell:
	UploadServer [port=2019] [path=.]

Example:
	UploadServer 2019 .
	UploadServer 8080 ~
	UploadServer 9090 static`

// return port and path
func initParameter() (int, string) {
	n := len(os.Args)
	if n <= 1 {
		return 2019, "."
	}
	port, err := strconv.Atoi(os.Args[1])
	if n == 2 {
		if err != nil {
			return 2019, os.Args[1]
		}
		return port, "."
	} else {
		if err != nil {
			port, err = strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(UploadServerRemark)
				os.Exit(2)
			}
			return port, os.Args[1]
		}
		return port, os.Args[2]
	}
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

func NewFileInfo(err error, size int64, httpname, filename string) map[string]interface{} {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return map[string]interface{}{
		"success":  err == nil,
		"message":  msg,
		"size":     size,
		"httpname": httpname,
		"filename": filename,
	}
}

func uploadHandler(dir string) http.HandlerFunc {
	static := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			static.ServeHTTP(w, r)
			return
		}
		err := r.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			static.ServeHTTP(w, r)
			return
		}
		if len(r.MultipartForm.File) == 0 {
			static.ServeHTTP(w, r)
			return
		}
		fis := make([] map[string]interface{}, 0)
		for name, file := range r.MultipartForm.File {
			for _, v := range file {
				f, err := os.OpenFile("."+r.URL.Path+v.Filename, os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					fis = append(fis, NewFileInfo(err, v.Size, name, v.Filename))
					continue
				}
				r, err := v.Open()
				if err != nil {
					fis = append(fis, NewFileInfo(err, v.Size, name, v.Filename))
					f.Close()
					continue
				}
				io.Copy(f, r)
				r.Close()
				f.Close()
				fis = append(fis, NewFileInfo(err, v.Size, name, v.Filename))
			}
		}
		w.Header().Set("Content-Type", "Content-Type:application/json")
		data, _ := json.Marshal(fis)
		w.Write(data)
	}
}

func main() {
	port, dir := initParameter()
	port_s := strconv.Itoa(port)
	fmt.Println("> Listening at http://127.0.0.1:" + port_s)
	http.HandleFunc("/", loggingHandler(uploadHandler(dir)))
	err := http.ListenAndServe(":"+port_s, nil)
	if err != nil {
		println(err.Error())
	}
}
