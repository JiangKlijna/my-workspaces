package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const ProxyServerRemark = `ProxyServer execute shell:
    ProxyServer [proxy port...] [local port] (iphash|random|round)

Example:
    ProxyServer 80 8080
	ProxyServer 8081,8082 8080 iphash
    ProxyServer 8080,192.168.1.1:8080 80 random
    ProxyServer 192.168.1.1:80,192.168.1.2:80 80 round`

func getPorts() (string, []string) {
	isNumber := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	if len(os.Args) > 2 {
		ports := make([]string, 0)
		arr := strings.Split(os.Args[2], ",")
		for _, p := range arr {
			if isNumber(p) {
				ports = append(ports, p)
			}
		}
		if len(arr) == 0 {
			return os.Args[1], []string{"8080", "8081", "8082"}
		} else {
			return os.Args[1], ports
		}
	}
	return "5050", []string{"8080", "8081", "8082"}
}

func getReverseProxies(ports []string) []http.Handler {
	hs := make([]http.Handler, len(ports))
	for i, port := range ports {
		targetUrl, _ := url.Parse("http://127.0.0.1:" + port)
		hs[i] = httputil.NewSingleHostReverseProxy(targetUrl)
	}
	return hs
}

// iphash
func clusterHandlerIphash(hs []http.Handler) http.Handler {
	n := uint32(len(hs))
	const primeRK = 16777619
	// 针对Remote Address所写的hash算法
	hashStr := func(str string) uint32 {
		hash := uint32(0)
		for _, v := range str {
			if v == ':' {
				return hash
			}
			hash = hash*primeRK + uint32(v)
		}
		return hash
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs[hashStr(r.RemoteAddr)%n].ServeHTTP(w, r)
	})
}

// random
func clusterHandlerRandom(hs []http.Handler) http.Handler {
	n := len(hs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hs[rand.Intn(n)].ServeHTTP(w, r)
	})
}

// round
func clusterHandlerRound(hs []http.Handler) http.Handler {
	n := len(hs)
	i := 0
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if i == n {
			i = 0
		}
		hs[i].ServeHTTP(w, r)
		i++
	})
}

func clusterHandler(hs []http.Handler) http.Handler {
	n := len(hs)
	if n == 1 {
		return hs[0]
	}
	fs := [3]func(hs []http.Handler) http.Handler{
		clusterHandlerRound, clusterHandlerRandom, clusterHandlerIphash,
	}
	return fs[rand.Intn(3)%3](hs)
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
		go fmt.Print(str)
	})
}

func main() {
	listeningPort, proxyPort := getPorts()
	fmt.Println("> Listening at http://127.0.0.1:" + listeningPort)
	http.Handle("/", loggingHandler(clusterHandler(getReverseProxies(proxyPort))))
	http.ListenAndServe(":"+listeningPort, nil)
}
