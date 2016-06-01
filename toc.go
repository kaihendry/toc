package toc

import (
	"bytes"
	"html/template"
	"io"

	"golang.org/x/net/html"
)

type header struct {
	Text string
	ID   string
}

func in(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// HTML finds the headers with ids and populates nav with them
func HTML(w io.Writer, r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	hx := []header{}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && in([]string{"h1", "h2", "h3", "h4", "h5", "h6"}, n.Data) {
			for _, a := range n.Attr {
				if a.Key == "id" {
					// fmt.Println(a.Val)
					hx = append(hx, header{Text: n.FirstChild.Data, ID: a.Val})
					break
				}
			}
			// fmt.Printf("%+v\n", n)
			// fmt.Println(n.FirstChild.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	// fmt.Println(hx)

	t, err := template.New("foo").Parse(`<ol>
{{- range . }}
<li><a href="#{{ .ID }}">{{ .Text }}</a></li>
{{- end }}
</ol>`)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, hx)
	// fmt.Println(buf.String())

	var insert func(*html.Node)
	insert = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "nav" {
			//fmt.Println("Found the NAV")
			//fmt.Printf("%+v\n", n)
			nodes, _ := html.ParseFragment(buf, n)
			for _, node := range nodes {
				n.AppendChild(node)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			insert(c)
		}
	}
	insert(doc)

	html.Render(w, doc)
	return nil

}
