package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s, welcome to your HTTP server🚀🚀.", r.URL.Path[1:])
}

func main() {
	fmt.Println("HTTP Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

// http://localhost:9000
// http://localhost:9000/text
