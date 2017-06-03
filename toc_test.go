package toc

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCreateTOC(t *testing.T) {
	var src, dst bytes.Buffer
	src.Write([]byte(`<h1>FAQ</h1><nav data-fill-with="table-of-contents" id="toc"></nav>
	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>`))
	Create(&dst, &src)

	output, _ := ioutil.ReadAll(&dst)
	if !strings.Contains(string(output), `<li><a href="#how-do-i-create-a-faq">How do I create a FAQ?</a></li>`) {
		t.Fail()
		return
	}
}
