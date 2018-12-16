//go build -ldflags "-s -w" HttpServer.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const StaticServerRemark = `StaticServer execute shell:
    StaticServer [port=80] [path=.]

Example:
	StaticServer 80 .
    StaticServer 8080 ~
	StaticServer 9090 static`

// return port and path
func initParameter() (int, string) {
	n := len(os.Args)
	if n <= 1 {
		fmt.Println(StaticServerRemark)
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if n == 2 {
		if err != nil {
			return 80, os.Args[1]
		}
		return port, "."
	} else {
		if err != nil {
			port, err = strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(StaticServerRemark)
				os.Exit(2)
			}
			return port, os.Args[1]
		}
		return port, os.Args[2]
	}
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		str := fmt.Sprintf(
			"%s Comleted %s %s in %v from %s\n",
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			time.Since(start),
			r.RemoteAddr)
		fmt.Print(str)
	})
}

func main() {
	port, path := initParameter()
	fmt.Println("> Listening at http://127.0.0.1:" + string(port))
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(path))))
	http.ListenAndServe(":"+string(port), nil)
}
