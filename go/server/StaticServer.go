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


func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getPort() string {
	if len(os.Args) > 1 && isNumber(os.Args[1]) {
		return os.Args[1]
	}
	return "5050"
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
	port := getPort()
	fmt.Println("> Listening at http://127.0.0.1:" + port)
	http.Handle("/", loggingHandler(http.FileServer(http.Dir("./"))))
	http.ListenAndServe(":" + port, nil)
}
