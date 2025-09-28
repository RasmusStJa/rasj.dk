package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	//"log"
	//"net/http"
)

var legalTypes = []string{"button", "collapse", "end"}

type link struct {
	baseURL string
	dir     []string
}

func (l *link) resetDir()            { l.dir = nil }
func (l *link) addDir(newDir string) { l.dir = append(l.dir, newDir) }
func (l *link) rmDirs(count int) {
	length := len(l.dir)
	if length == count {
		l.resetDir()
	} else if count > 0 && length > count {
		l.dir = l.dir[:length-count]
	}
}

type navBtn struct {
	name, dir, btnType string
}

func (btn *navBtn) setName(name string)    { btn.name = name }
func (btn *navBtn) setDir(dir string)      { btn.dir = dir }
func (btn *navBtn) setType(btnType string) { btn.btnType = btnType }
func (btn navBtn) HTML(url *link) {
	switch btn.btnType {
	case legalTypes[0]: //button
		fmt.Println("<button>", btn.name, "</button>")
	case legalTypes[1]: //collapse
		fmt.Println("<button id='toggleButton'>", btn.name, ":", "</button>")
		fmt.Println("<div id='collapsibleDiv' style='display: none;'>")
		url.addDir(btn.dir)
	case legalTypes[2]: //end
		//TODO: come up with better name than "name" (should be smth like amount of ends)
		switch strings.ToLower(btn.name) {
		case "all": //remove everything
			for range len(url.dir) {
				fmt.Println("</div>")
			}
			url.resetDir()
		case "": //remove 1 time
			fmt.Println("</div>")
			url.rmDirs(1)
		default: //remove n times
			count, err := strconv.Atoi(btn.name)
			if err != nil {
				panic(err)
			}
			for range count {
				fmt.Println("</div>")
			}
			url.resetDir()
		}
	}

}

type navBar struct{ btnContainer []navBtn }

func (nb navBar) appendBtn(btn navBtn) {
	nb.btnContainer = append(nb.btnContainer, btn)
}
func (nb navBar) HTML() {
	url := link{baseURL: "https://rasj.dk"}
	for _, btn2 := range nb.btnContainer {
		if slices.Contains(legalTypes, btn2.btnType) {
			btn2.HTML(&url)
			continue
		}
		fmt.Println(btn2.name, "has illegal type:", btn2.btnType)
		break
	}
}

func main() {
	var mainNavBar navBar
	mainNavBar.appendBtn(navBtn{name: "Forside", dir: "/", btnType: "button"})
	mainNavBar.appendBtn(navBtn{name: "Spil", dir: "/spil", btnType: "collapse"})
	mainNavBar.appendBtn(navBtn{name: "Wordle", dir: "/wordle", btnType: "button"})
	mainNavBar.appendBtn(navBtn{btnType: "end"})

	mainNavBar.HTML()
}
