package main

import (
	"runtime"
	"strings"
)

type favicon struct {
	ftype string
	fpath string
}
type header struct {
	fv        favicon
	stylepath string
}

func (h header) HTML() string {
	head := element{tag:"header"}

	stylesheet := element{tag: "link"}
	stylesheet.attributes = []attribute{{name: "href", value: "https://rasj.dk" + h.stylepath}, {name: "rel", value: "stylesheet"}}

	charset := element{tag: "meta"}
	charset.attributes = []attribute{{name: "charset", value: "utf-8"}}

	
	viewport := element{tag: "meta"}
	viewport.attributes = []attribute{{name: "name", value: "viewport"}, {name: "content", value: "width=device-width, initial-scale=1"}}

	icon := element{tag: "link"}
	icon.attributes = []attribute{{name: "rel", value: "icon"}, {name: "type", value: "image/svg+xml"}, {name: "href", value: "https://rasj.dk" + h.fv.fpath}}

	mathjax := element{tag: "script"}
	mathjax.attributes = []attribute{{name: "id", value: "MathJax-script"}, {name: "async"}, {name: "src", value: "https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"}}

	_, filename, _, _ := runtime.Caller(1)
	title := element{tag: "title"}
	title.innerText = "pi.rasj.dk - " + strings.Split(filename, ".")[0]

	head.AppendChild(stylesheet)
	head.AppendChild(charset)
	head.AppendChild(viewport)
	head.AppendChild(icon)
	head.AppendChild(mathjax)
	head.AppendChild(title)

	return head.HTML()
}
