package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown"
)

func main() {
	file := "test.md"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found!", file)
	}

	md := []byte(content)
	html := markdown.ToHTML(md, nil, nil)

	fmt.Print(string(html))
	// fmt.Printf(string(html))

	fileWrite := "test.html"
	writeErr := ioutil.WriteFile(fileWrite, html, 0644)

	if writeErr != nil {
		log.Fatalf("Could not write to the file %s", fileWrite)
	}

	fmt.Printf("See the HTML output in %s file", fileWrite)
}
