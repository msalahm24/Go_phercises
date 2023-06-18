package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type link struct {
	href string
	text string
}

func main() {
	flagHTMLFileName := flag.String("html", "ex1.html", "the html file to be parsed")
	flag.Parse()

	f, err := os.Open(*flagHTMLFileName)
	if err != nil {
		log.Fatalf("Failed to open the html file %q: %v", *flagHTMLFileName, err)
	}

	defer f.Close()
	root, err := html.Parse(f)
	if err != nil {
		log.Fatalf("Failed to parse the html file %q: %v", *flagHTMLFileName, err)
	}

	as := make(chan *html.Node)
	go findAnchors(root, as)

	for a := range as {
		fmt.Println(link{
			href: extractHref(a),
			text: extractText(a),
		})
	}
}

func findAnchors(n *html.Node, as chan *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		as <- n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchors(c, as)
	}

	if n.Parent == nil {
		close(as)
	}
}

func extractHref(a *html.Node) string {
	for _,attr := range a.Attr{
		if attr.Key != "href"{
			continue
		}
		return attr.Val
	}
	return ""
}

func extractText(a *html.Node) string{
	var text string
	for c := a.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode{
			text += c.Data
		}
		text += extractText(c)
	}
	return strings.TrimSpace(text) 
}
