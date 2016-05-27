package main

import (
	"bytes"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	h, err := os.Open("test.src.html")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(h)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "nav" {
			// fmt.Printf("%+v\n", n)
			nodes, _ := html.ParseFragment(bytes.NewBufferString("<p>howdy</p>"), n)
			for _, node := range nodes {
				n.AppendChild(node)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	html.Render(os.Stdout, doc)
}
