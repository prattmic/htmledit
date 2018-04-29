package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func fixBody(d *goquery.Document) {
	// For some reason, the parser adds an extra newline to the text just
	// before </body>. Drop it.
	s := d.Find("body")
	if s.Length() == 0 {
		fmt.Fprintf(os.Stderr, "body not found\n")
		return
	}

	n := s.Get(0).LastChild
	if n == nil {
		fmt.Fprintf(os.Stderr, "body has no children\n")
		return
	}

	if n.Type != html.TextNode {
		fmt.Fprintf(os.Stderr, "body last child %+v not text\n", n)
		return
	}

	if len(n.Data) < 1 {
		fmt.Fprintf(os.Stderr, "body last child has no data\n")
		return
	}

	if n.Data[len(n.Data)-1] != '\n' {
		fmt.Fprintf(os.Stderr, "body last child does not end in newline: %q\n", n.Data)
		return
	}

	n.Data = n.Data[:len(n.Data)-1]
}

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

	fixBody(d)

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
