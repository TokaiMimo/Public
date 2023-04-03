package main

import (
	"fmt"
	"net/http"
	"strings"
)

// type MyHandler struct {
// }
//
//	func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		if r.URL.Path == "/" {
//			w.Write([]byte("Home"))
//			return
//		}
//		if strings.HasPrefix(r.URL.Path, "/hello") {
//			name := strings.Split(r.URL.Path, "/")[2]
//			w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
//			return
//		}
//
//		w.WriteHeader((http.StatusNotFound))
//		w.Write([]byte("404 Page Not Found"))
//	}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)

	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	return
}
