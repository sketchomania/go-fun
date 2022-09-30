package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	fmt.Println("Welcome to text to pdf converter")
	file := "test.txt"
	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(30)
	pdf.AddPage()
	pdf.SetFont("Times", "B", 12)

	pdf.MultiCell(190, 5, string(content), "0", "0", false)

	_ = pdf.OutputFileAndClose("generated.pdf")

	fmt.Println("PDF Created")
}
