package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"strconv"
	"os"
)

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

func templateHandler(root http.FileSystem) http.HandlerFunc {
	next := http.FileServer(root)
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.RequestURI, ".html") {
			templateFile(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func templateFile(w http.ResponseWriter, r *http.Request) {
	upath := strings.TrimPrefix(r.URL.Path, "/")

	t, err := template.ParseFiles(upath)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	r.ParseForm()
	t.Execute(w, map[string]interface{}{
		"addr":    r.RemoteAddr,
		"cookies": r.Cookies(),
		"form":    r.Form,
		"header":  r.Header,
		"uri":     r.RequestURI,
	})
}

func server(port string) {
	fmt.Println("> Listening at http://127.0.0.1:" + port)
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(templateHandler(http.Dir("./"))))
	http.ListenAndServe(":"+port, mux)
}

func main() {
	server(getPort())
}
