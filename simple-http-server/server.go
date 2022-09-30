package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func loadFile(fileName string) (string, error) {
	bytes_data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes_data), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%v", r.URL)
	// fmt.Printf("%v", path.Base(r.URL.EscapedPath()))
	var index, err = loadFile(path.Base(r.URL.EscapedPath()))
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404: Oops! Couldn't find the page you are looking for...")
	}
	fmt.Fprint(w, index)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

// http://localhost:9000
// http://localhost:9000/index.html
// http://localhost:9000/ere/todo/index.html
