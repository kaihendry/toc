package main

import (
	"fmt"
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
		panic(err)
	}
	defer html.Close()

	toc.CreateTOC(os.Stdout, html)

}
