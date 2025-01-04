package main

import (
	"fmt"
	"log"
	"net/http"
)

// start command: go run local_file_server.go
func main() {
	// root "/"
	http.HandleFunc("/sketchomania", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sketchomania" {
			// Return 404 for any unhandled paths
			http.NotFound(w, r)
			return
		}

		fmt.Fprintln(w, "Welcome to the File Server!\nAccess shared folders at: /dir1, /dir2, /dir3, /to-edit, /imp. \n Modify access if not found")
	})

	dir1 := http.FileServer(http.Dir("D:/dir1"))
	http.Handle("/sketchomania/dir1/", http.StripPrefix("/sketchomania/dir1/", dir1))

	// dir2 := http.FileServer(http.Dir("D:/dir2"))
	// http.Handle("/sketchomania/dir2/", http.StripPrefix("/sketchomania/dir2/", dir2))

	// dir3 := http.FileServer(http.Dir("D:/dir3"))
	// http.Handle("/sketchomania/dir3/", http.StripPrefix("/sketchomania/dir3/", dir3))

	log.Println("Serving files on http://localhost:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
