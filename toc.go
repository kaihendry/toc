package toc

import (
	"bytes"
	"html/template"
	"io"

	"golang.org/x/net/html"
)

var t = template.Must(template.New("tocheadermarkup").Parse(`<ol>
{{- range . }}
<li><a href="#{{ .ID }}">{{ .Text }}</a></li>
{{- end }}
</ol>`))

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

// CreateTOC creates the table of contents by finding the headers and inserting them as an ordered list in the placeholder
func CreateTOC(dst io.Writer, src io.Reader) error {
	doc, err := html.Parse(src)
	if err != nil {
		return err
	}

	headers := []header{}
	getHeaders(&headers, doc)
	// fmt.Println(headers)

	buf := new(bytes.Buffer)
	t.Execute(buf, headers)

	insertTOCNodes(buf, doc)
	html.Render(dst, doc)

	return nil
}
