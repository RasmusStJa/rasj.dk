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

func (e element) HTML() string {
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
