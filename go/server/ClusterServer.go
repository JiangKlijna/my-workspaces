//go build -ldflags "-s -w" ClusterServer.go
package main

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"strconv"
)

var ico = []byte{
	137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 16, 0, 0, 0, 16, 8, 6, 0, 0, 0, 31, 243, 255, 97, 0, 0, 0, 51, 73, 68, 65, 84, 120, 156, 99, 248, 240, 225, 195, 127, 74, 48, 3, 186, 192, 196, 235, 115, 225, 24, 159, 216, 32, 54, 128, 226, 48, 24, 169, 6, 32, 135, 50, 41, 120, 144, 26, 128, 45, 81, 225, 75, 104, 131, 196, 0, 74, 48, 0, 42, 31, 140, 146, 20, 142, 145, 240, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130,
}

var html [12][]byte = [12][]byte{
	[]byte("<!DOCTYPE html><html><head><title>Server"),
	[]byte("</title><style>*{font-family:consolas;}</style></head>"),
	[]byte("<body><table><tr><th>General</th><th></th></tr>"),
	[]byte("<tr><td>Listening Port</td><td>" ),
	[]byte("</td></tr><tr><td>Request URL</td><td>" ),
	[]byte("</td></tr><tr><td>Request Method</td><td>" ),
	[]byte("</td></tr><tr><td>Remote Address</td><td>" ),
	[]byte("</td></tr></table><table><tr><th>Request Headers</th><th></th></tr>"),
	[]byte("<tr><td>"), []byte("</td><td>"), []byte("</td></tr>"),
	[]byte("</table></body></html>"),
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getPorts() []string {
	if len(os.Args) <= 1 {
		return []string{"8080", "8081", "8082"};
	}
	ports := make([]string, 0)
	for i := 1; i < len(os.Args); i++ {
		if isNumber(os.Args[i]) {
			ports = append(ports, os.Args[i])
		}
	}
	return ports
}

func singleHandler(port string) http.Handler {
	outputHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			w.Write(ico)
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.Write(html[0])
			w.Write([]byte(port))
			w.Write(html[1])
			w.Write(html[2])
			w.Write(html[3])
			w.Write([]byte(port))
			w.Write(html[4])
			w.Write([]byte(r.RequestURI))
			w.Write(html[5])
			w.Write([]byte(r.Method))
			w.Write(html[6])
			w.Write([]byte(r.RemoteAddr))
			w.Write(html[7])
			for key, value := range r.Header {
				w.Write(html[8])
				w.Write([]byte(key))
				w.Write(html[9])
				if len(value) == 1 {
					w.Write([]byte(value[0]))
				} else {
					w.Write([]byte(fmt.Sprint(value)))
				}
				w.Write(html[10])
			}
			w.Write(html[11])
		}
	}
	loggingHandler := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		outputHandler(w, r)
		str := fmt.Sprintf(
			"Server %s: %s Comleted %s %s in %v from %s\n",
			port, start.Format("2006-01-02 15:04:05"),
			r.Method, r.URL.Path, time.Since(start), r.RemoteAddr)
		fmt.Print(str)
	}
	return http.HandlerFunc(loggingHandler)
}

func server(port string) {
	fmt.Println("> Listening at http://127.0.0.1:" + port)
	mux := http.NewServeMux()
	mux.Handle("/", singleHandler(port))
	http.ListenAndServe(":"+port, mux)
}

func main() {
	done := make(chan bool)
	ports := getPorts()
	for _, port := range ports {
		go server(port)
	}
	<-done
}
