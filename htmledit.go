package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s file.html\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	d, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		fmt.Printf("error parsing file: %v\n", err)
		os.Exit(1)
	}

	h, err := goquery.OuterHtml(d.Selection)
	if err != nil {
		fmt.Printf("error generating html: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", h)
}
