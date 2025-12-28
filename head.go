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
	var head, stylesheet, charset, viewport, icon, mathjax, title element
	head.tag = "header"

	stylesheet.tag = "link"
	stylesheet.AppendAttribute(attribute{name: "href", value: "https://rasj.dk" + h.stylepath})
	stylesheet.AppendAttribute(attribute{name: "rel", value: "stylesheet"})

	charset.tag = "meta"
	charset.AppendAttribute(attribute{name: "charset", value: "utf-8"})

	viewport.tag = "meta"
	viewport.AppendAttribute(attribute{name: "name", value: "viewport"})
	viewport.AppendAttribute(attribute{name: "content", value: "width=device-width, initial-scale=1"})

	icon.tag = "link"
	icon.AppendAttribute(attribute{name: "rel", value: "icon"})
	icon.AppendAttribute(attribute{name: "type", value: "image/svg+xml"})
	icon.AppendAttribute(attribute{name: "href", value: "https://rasj.dk" + h.fv.fpath})

	mathjax.tag = "script"
	mathjax.AppendAttribute(attribute{name: "id", value: "MathJax-script"})
	mathjax.AppendAttribute(attribute{name: "async"})
	mathjax.AppendAttribute(attribute{name: "src", value: "https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"})

	_, filename, _, _ := runtime.Caller(1)
	title.tag = "title"
	title.innerText = "rasj.dk - " + strings.Split(filename, ".")[0]

	head.AppendChild(stylesheet)
	head.AppendChild(charset)
	head.AppendChild(viewport)
	head.AppendChild(icon)
	head.AppendChild(mathjax)
	head.AppendChild(title)

	return head.HTML()
}
