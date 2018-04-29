package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var selector string
	switch len(os.Args) {
	case 1:
	case 2:
		selector = os.Args[1]
	default:
		fmt.Fprintf(os.Stderr, "usage: %s [selector]\n", os.Args[0])
		os.Exit(1)
	}

	d, err := goquery.NewDocumentFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing file: %v\n", err)
		os.Exit(1)
	}

	if selector != "" {
		s := d.Find(selector)
		if s.Length() == 0 {
			fmt.Fprintf(os.Stderr, "selector %q not found\n", selector)
			os.Exit(1)
		}
		s.Remove()
	}

	h, err := goquery.OuterHtml(d.Selection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error generating html: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", h)
}
