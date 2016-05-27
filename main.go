package main

import (
	"fmt"
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
		if n.Type == html.ElementNode && n.Data == "h3" {
			fmt.Printf("%+v\n", n)
			fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
