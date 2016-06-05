package toc

import (
	"bytes"
	"fmt"
	"html/template"
	"io"

	"golang.org/x/net/html"
)

var (
	headerTags = []string{"h1", "h2", "h3", "h4", "h5", "h6"}
)

type header struct {
	Text string
	ID   string
}

func isHeader(node *html.Node) bool {
	if node.Type != html.ElementNode {
		return false
	}

	for _, tag := range headerTags {
		if tag == node.Data {
			return true
		}
	}

	return false
}

func getHeaders(headers *[]header, node *html.Node) {
	if isHeader(node) {
		for _, attr := range node.Attr {
			if attr.Key == "id" {
				*headers = append(*headers, header{node.FirstChild.Data, attr.Val})
				break
			}
		}
	}

	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		getHeaders(headers, nextNode)
	}
}

func isTOCPlaceholder(node *html.Node) bool {
	if node.Type == html.ElementNode {
		for _, attr := range node.Attr {
			if attr.Key == "data-fill-with" &&
				attr.Val == "table-of-contents" {
				return true
			}
		}
	}
	return false
}

func insertTOCNodes(buf *bytes.Buffer, node *html.Node) {
	if isTOCPlaceholder(node) {
		nodes, _ := html.ParseFragment(buf, node)
		for _, n := range nodes {
			node.AppendChild(n)
		}
	}

	for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		insertTOCNodes(buf, nextNode)
	}
}

// Insert does what it says.. It inserts stuff into other stuff. (please write something nice here)
func Insert(dst io.Writer, src io.Reader) error {
	doc, err := html.Parse(src)
	if err != nil {
		return err
	}

	headers := []header{}
	getHeaders(&headers, doc)
	fmt.Println(headers)

	// TODO: make this nicer?
	t, _ := template.New("foo").Parse(`<ol>
{{- range . }}
<li><a href="#{{ .ID }}">{{ .Text }}</a></li>
{{- end }}
</ol>`)

	buf := new(bytes.Buffer)
	t.Execute(buf, headers)

	insertTOCNodes(buf, doc)
	html.Render(dst, doc)

	return nil
}
