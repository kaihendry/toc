# Creates a table of contents from [h1, h2, h3, h4, h5, and h6 elements](https://html.spec.whatwg.org/multipage/semantics.html#the-h1,-h2,-h3,-h4,-h5,-and-h6-elements)

[![Build Status](https://travis-ci.org/kaihendry/toc.svg?branch=master)](https://travis-ci.org/kaihendry/toc)

<img src=http://s.natalian.org/2016-06-05/toc.gif alt="Demo">

An element with `data-fill-with="table-of-contents"` is filled with an ordered list of the headers.

This allows users to hyperlink directly to answers.

## Install

To install:

	go install github.com/kaihendry/toc/cmd/toc

To keep upto date:

	go get -u github.com/kaihendry/toc/cmd/toc

## Example

`foo.html` contains:

	<h1>FAQ</h1>
	<nav data-fill-with="table-of-contents" id="toc"></nav>
	<h3 id="how-do-i-create-a-faq">How do I create a FAQ?</h3><p>Like this!</p>

Run the tool over the HTML

	$ toc foo.html

And something like the following should print to `/dev/stdout`

	<h1>FAQ</h1>
	<nav data-fill-with=table-of-contents id=toc>
	<ol>
	<li><a href=#how-do-i-create-a-faq>How do I create a FAQ?</a>
	</ol>
	</nav>
	<h3 id=how-do-i-create-a-faq>How do I create a FAQ?</h3>
	<p>Like this!</p>

As used on:

* <https://webconverger.com/faq/>

## Acknowledgements

Many thanks to [tabatkins's bikeshed](https://github.com/tabatkins/bikeshed) for the inspiration & the most helpful [Golang programmer I know, Zoltan](https://github.com/zgiber).
