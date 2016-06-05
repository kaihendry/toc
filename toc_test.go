package toc

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestInsert(t *testing.T) {
	var src, dst bytes.Buffer
	src.Write([]byte(`<h1>FAQ</h1><nav data-fill-with="table-of-contents" id="toc"></nav>
	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>`))
	Insert(&dst, &src)

	// @Kai this is kind of naive and limited, you probably want to improve the tests
	output, _ := ioutil.ReadAll(&dst)
	if !strings.Contains(string(output), `<li><a href="#how-do-i-create-a-faq">How do I create a FAQ?</a></li>`) {
		t.Fail()
		return
	}
	// Normalise HTML for my sanity, though the minifier function could change and make the output incorrect
	// gosml.HTML(os.Stdout, &b2)
	// Output: <h1>FAQ</h1><nav data-fill-with=table-of-contents id=toc><ol><li><a href=#how-do-i-create-a-faq>How do I create a FAQ?</a></ol></nav><h3 id=how-do-i-create-a-faq>How do I create a FAQ?</h3><p>Like this!
}

// Do you want to keep this canonicalise stuff? Why is it good? :)
// func ExampleCanonicalisation() {
// 	var b bytes.Buffer
// 	b.Write([]byte(`<h1>	Hello     World
// 	</h1>`))
// 	canonicalise(os.Stdout, &b)
// 	// Output: <h1>Hello World</h1>
// }
