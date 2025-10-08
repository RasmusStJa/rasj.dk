package main

type attribute struct {
	name  string
	value string
}

func (a attribute) HTML() string {
	if a.value != "" {
		return " " + a.name + "='" + a.value + "' "
	}
	return " " + a.name + " "
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
	e.attributes = []attribute{{name: "onclick", value: "location.href=\"" + url + "\""}}
}

func (e *element) CreateBody() {
	e.tag = "body"
	e.attributes = []attribute{{name: "style", value: "font-family:monospace;"}}
}

func (e *element) CreateNavBar() {
	e.tag = "div"
	btn := element{}
	btn.CreateBtn("About", "https://pi.rasj.dk/about")
	e.AppendChild(btn)
	btn.CreateBtn("Is now a prime?", "https://pi.rasj.dk/isnowaprime")
	e.AppendChild(btn)
}

func (e element) HTML() string {
	if e.tag == "" {
		panic("This element has an undefined tag: " + e.HTML())
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
