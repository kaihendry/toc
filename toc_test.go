package main

import (
	"bytes"
	"os"

	"github.com/kaihendry/gosml"
)

func ExampleTOC() {
	var b, b2 bytes.Buffer
	b.Write([]byte(`<h1>FAQ</h1><nav data-fill-with="table-of-contents" id="toc"></nav>
	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>`))
	toc(&b2, &b)
	// Normalise HTML for my sanity, though the minifier function could change and make the output incorrect
	gosml.HTML(os.Stdout, &b2)
	// Output: <h1>FAQ</h1><nav data-fill-with=table-of-contents id=toc><ol><li><a href=#how-do-i-create-a-faq>How do I create a FAQ?</a></ol></nav><h3 id=how-do-i-create-a-faq>How do I create a FAQ?</h3><p>Like this!
}

func ExampleCanonicalisation() {
	var b bytes.Buffer
	b.Write([]byte(`<h1>	Hello     World
	</h1>`))
	canonicalise(os.Stdout, &b)
	// Output: <h1>Hello World</h1>
}
