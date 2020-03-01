package main

// curl http://localhost:8080/
// URL.Path = "/"
// curl http://localhost:8080/hello
// Header["User-Agent"] = ["curl/7.65.1"]
// Header["Accept"] = ["*/*"]

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("start http server")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
