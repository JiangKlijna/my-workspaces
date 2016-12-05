package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func httpUserInfo(w http.ResponseWriter, r *http.Request) {
	ret := "your addr is:" + r.RemoteAddr + "\r\n"
	ret += "request headers:" + fmt.Sprint(r.Header)
	io.WriteString(w, ret)
}

func main() {
	targetURL, err := url.Parse("http://127.0.0.1:9090")
	log.Println("program start......")
	if err != nil {
		log.Println("bad url:", err)
		return
	}
	httpProxy := httputil.NewSingleHostReverseProxy(targetURL)
	http.HandleFunc("/myinfo/", httpUserInfo)
	http.Handle("/", httpProxy)
	log.Println("listen http proxy on :5050.....")
	http.ListenAndServe(":5050", nil)
}