package main

import (
	"fmt"
	"slices"
	"strconv"
	//"log"
	//"net/http"
)

var legalTypes = []string{"button", "collapse", "end"}

type link struct {
	baseURL string
	dir     []string
}

func (l link) resetDir()            { l.dir = nil }
func (l link) addDir(newDir string) { l.dir = append(l.dir, newDir) }
func (l link) rmDirs(count int) {
	length := len(l.dir)
	if length == count {
		l.resetDir()
	} else if length > 0 && count > 0 {
		l.dir = l.dir[:length-count]
	}
}

type navBtn struct {
	name, dir, btnType string
}

func (btn navBtn) setName(name string)    { btn.name = name }
func (btn navBtn) setDir(dir string)      { btn.dir = dir }
func (btn navBtn) setType(btnType string) { btn.btnType = btnType }
func (btn navBtn) html(url link) {
	switch btn.btnType {
	case legalTypes[0]: //button
		fmt.Println("<button>", "</button>")
	case legalTypes[1]: //collapse
		fmt.Println("<button id='toggleButton'>", btn.name, ":", "</button> <div id='collapsibleDiv' style='display: none;'>")
	case legalTypes[2]: //end
		strconv.Atoi(btn.name)
	}

}

type navBar struct{ btnContainer []navBtn }

func (nb navBar) appendBtn(btn navBtn) {
	nb.btnContainer = append(nb.btnContainer, btn)
}
func (nb navBar) html() {
	url := link{baseURL: "https://rasj.dk"}
	for _, btn2 := range nb.btnContainer {
		if slices.Contains(legalTypes, btn2.btnType) {
			btn2.html(url)
		} else {
			fmt.Println(btn2.name, "has illegal type:", btn2.btnType)
			break
		}
	}
}

func main() {
	var mainNavBar navBar
	mainNavBar.appendBtn(navBtn{name: "Forside", dir: "/", btnType: "button"})
	fmt.Println("Hello World!")
	mainNavBar.html()
}
