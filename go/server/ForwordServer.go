//go build -ldflags "-s -w" ForwordServer.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ForwordServerRemark = `ForwordServer execute shell:
	ForwordServer [port=2019]

Example:
	ForwordServer 2019`

func getPort() string {
	if len(os.Args) > 1 {
		_, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(ForwordServerRemark)
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

var client = &http.Client{Transport: &http.Transport{Proxy: nil}}

func forwordHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery == "" {
		http.NotFound(w, r)
		return
	}
	url := strings.TrimPrefix(strings.Join(strings.Split(r.URL.RawQuery, "url="), "&"), "&")
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for k := range r.Header {
		req.Header.Add(k, r.Header.Get(k))
	}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for k := range res.Header {
		w.Header().Add(k, res.Header.Get(k))
	}
	io.Copy(w, res.Body)
}

func main() {
	port := getPort()
	fmt.Println("> Listening at http://127.0.0.1:" + port)
	http.HandleFunc("/forword", loggingHandler(forwordHandler))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		println(err.Error())
	}
}
