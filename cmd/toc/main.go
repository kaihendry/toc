package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kaihendry/toc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("HTML file required")
		os.Exit(1)
	}
	html, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer html.Close()

	toc.Create(os.Stdout, html)

}
