package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadFile(fileName string) (string, error) {
	bytes_data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes_data), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var index, err = loadFile("index2.html")
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404: Oops! Couldn't find the page you are looking for...")
	}
	fmt.Fprint(w, index)
	// fmt.Fprintf(w, "Hello %s, welcome to your HTTP server🚀🚀.", r.URL.Path[1:])
}

func main() {
	fmt.Println("HTTP Server")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

// http://localhost:9000
// http://localhost:9000/text
