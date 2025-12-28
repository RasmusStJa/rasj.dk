package main

type attribute struct {
	name  string
	value string
}

func (a attribute) HTML() string {
	if a.value == "" {
		return " " + a.name + " "
	}
	return " " + a.name + "='" + a.value + "' "
}

type element struct {
	tag        string
	attributes []attribute
	innerText  string
	children   []element
}

func (e *element) AppendAttribute(attr attribute) {
	e.attributes = append(e.attributes, attr)
}

func (e *element) AppendChild(child element) {
	e.children = append(e.children, child)
}

func (e *element) AppendChildren(children []element) {
	for _, child := range children {
		e.AppendChild(child)
	}
}

func (e *element) Clear() {
	e.tag = ""
	e.attributes = []attribute{}
	e.innerText = ""
	e.children = []element{}
}

func (e *element) CreateBtn(name string, url string) {
	e.tag = "button"
	e.innerText = name
	e.AppendAttribute(attribute{name: "onclick", value: "location.href=\"https://" + url + "\""})
}

func (e *element) CreateLink(text string, domain string) {
	e.tag = "a"

	if text == "" {
		e.innerText = domain
	} else {
		e.innerText = text
	}

	e.AppendAttribute(attribute{name: "href", value: "https://" + domain})
}

func (e *element) CreateBody() {
	e.tag = "body"
	e.AppendAttribute(attribute{name: "style", value: "font-family:monospace;"})
}

func (e *element) CreateNavBar() {
	const url string = "rasj.dk/"
	var btn element

	e.tag = "div"

	btn.CreateBtn("About", url+"about")
	e.AppendChild(btn)

	btn.Clear()
	btn.CreateBtn("Is now a prime?", url+"isnowaprime")
	e.AppendChild(btn)

	btn.Clear()
	btn.CreateBtn("Er det fredag idag?", url+"fredag")
	e.AppendChild(btn)
}

func (e *element) HTML() string {
	if e.tag == "" {
		panic("This element has an undefined tag: " + e.innerText)
	}

	var result string

	result = "<" + e.tag
	for _, attr := range e.attributes {
		result += attr.HTML()
	}
	result += ">" + e.innerText

	for _, child := range e.children {
		result += child.HTML()
	}

	result += "</" + e.tag + ">"
	return result
}
