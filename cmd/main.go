package main

import (
	"fmt"
	"os"

	"github.com/zgiber/toc" // change this to your repo after merging (if you merge)
)

// // I'm not dure you want to keep this, so left it commented.. :)

// HTML Normalise/Standardise/Canonicalize HTML
// func canonicalise(w io.Writer, r io.Reader) error {

// 	var b bytes.Buffer
// 	var minifier = minify.New()
// 	minifier.AddFunc("text/html", minhtml.Minify)

// 	doc, err := html.Parse(r)
// 	if err != nil {
// 		return err
// 	}
// 	html.Render(&b, doc)

// 	m := minify.New()
// 	minhtml.Minify(m, w, &b, nil)

// 	return nil
// }

func main() {
	if len(os.Args) < 2 {
		fmt.Println("HTML file required")
		os.Exit(1)
	}
	html, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer html.Close()

	// var b bytes.Buffer
	toc.Insert(os.Stdout, html)
	// canonicalise(os.Stdout, &b)

}
