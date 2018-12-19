//go build -ldflags "-s -w" -o ps.exe ProxyServer.go
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
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

var re = regexp.MustCompile("^(\\d+|\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d+|http.+)$")

type ClusterHandler func([]http.Handler) http.Handler

func initParameter() ([]string, string, ClusterHandler) {
	if len(os.Args) < 3 {
		fmt.Print(ProxyServerRemark)
		os.Exit(1)
	}
	proxyServers := strings.Split(os.Args[1], ",")
	for _, s := range proxyServers {
		if !re.MatchString(s) {
			fmt.Print(ProxyServerRemark)
			os.Exit(2)
		}
	}
	isNumber := func(s string) bool {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	listeningPort := os.Args[2]
	if !isNumber(listeningPort) {
		fmt.Print(ProxyServerRemark)
		os.Exit(3)
	}
	var clusterHandler func([]http.Handler) http.Handler
	if len(proxyServers) == 1 {
		clusterHandler = clusterHandlerSingle
	} else if len(os.Args) == 3 {
		clusterHandler = clusterHandlerRound
	} else if len(os.Args) >= 4 {
		switch os.Args[3] {
		case "iphash":
			clusterHandler = clusterHandlerIphash
			break
		case "random":
			clusterHandler = clusterHandlerRandom
			break
		case "round":
			clusterHandler = clusterHandlerRound
			break
		default:
			fmt.Print(ProxyServerRemark)
			os.Exit(4)
		}
	}
	return proxyServers, listeningPort, clusterHandler
}

func getReverseProxies(proxyServers []string) []http.Handler {
	hs := make([]http.Handler, len(proxyServers))
	for i, server := range proxyServers {
		var targetUrl *url.URL
		var err error
		if strings.HasPrefix(server, "http") {
			targetUrl, err = url.Parse(server)
		} else if strings.Contains(server, ":") {
			targetUrl, err = url.Parse("http://" + server)
		} else {
			targetUrl, err = url.Parse("http://127.0.0.1:" + server)
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
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

//single
func clusterHandlerSingle(hs []http.Handler) http.Handler {
	return hs[0]
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
	proxyServers, listeningPort, clusterHandler := initParameter()
	fmt.Println("> Listening at http://127.0.0.1:" + listeningPort)
	http.Handle("/", loggingHandler(clusterHandler(getReverseProxies(proxyServers))))
	http.ListenAndServe(":"+listeningPort, nil)
}
