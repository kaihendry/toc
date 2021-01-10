package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kaihendry/toc"
)

const Version = "toc 0.0.3"

func usage() {
	fmt.Fprintf(os.Stderr, Version+`

HTML input with the attribute:

	data-fill-with="table-of-contents"

Will be filled with an ordered list of the headers.
`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		toc.Create(os.Stdout, os.Stdin)
		return
	}

	html, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer html.Close()

	toc.Create(os.Stdout, html)
}
