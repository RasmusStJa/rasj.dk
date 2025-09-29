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
	dirs    []string
}

func (l *link) resetDir()            { l.dirs = nil }
func (l *link) addDir(newDir string) { l.dirs = append(l.dirs, newDir) }
func (l *link) rmDirs(count int) {
	length := len(l.dirs)
	if length == count {
		l.resetDir()
	} else if count > 0 && length > count {
		l.dirs = l.dirs[:length-count]
	}
}
func (l link) get() string {
	result := l.baseURL
	for _, dir := range l.dirs {
		result += dir
	}
	return result
}

type navBtn struct {
	name, dir, btnType string
}

func (btn *navBtn) setName(name string)    { btn.name = name }
func (btn *navBtn) setDir(dir string)      { btn.dir = dir }
func (btn *navBtn) setType(btnType string) { btn.btnType = btnType }
func (btn navBtn) HTML(url *link) string {
	var result string
	switch btn.btnType {
	case legalTypes[0]: //button
		result = "<button href='" + url.get() + btn.dir + "'>" + btn.name + "</button>"
	case legalTypes[1]: //collapse
		url.addDir(btn.dir)
		result = "<button id='toggleButton'>" + btn.name + ":" + "</button>" + "<div id='collapsibleDiv' style='display: none;'>"
	case legalTypes[2]: //end
		//TODO: come up with better name than "name" (should be smth like amount of ends)
		switch strings.ToLower(btn.name) {
		case "all": //remove everything
			for range len(url.dirs) {
				result += "</div>"
			}

			url.resetDir()
		case "": //remove 1 time
			if len(url.dirs) >= 1 {
				url.rmDirs(1)
				result = "</div>"
			}
		default: //remove n times
			count, err := strconv.Atoi(btn.name) //string (anything) to int
			if err != nil {
				panic(err)
			}
			if len(url.dirs) >= count {
				for range count {
					result += "</div>"
				}
				url.resetDir()
			}
		}
	}
	return result
}

type navBar struct{ btnContainer []navBtn }

func (nb *navBar) appendBtn(btn navBtn) {
	nb.btnContainer = append(nb.btnContainer, btn)
}
func (nb navBar) HTML() string {
	var result string
	url := link{baseURL: "https://rasj.dk"}
	for _, btn2 := range nb.btnContainer {
		if slices.Contains(legalTypes, btn2.btnType) {
			result += btn2.HTML(&url)
			continue
		}
		fmt.Println(btn2.name, "has illegal type:", btn2.btnType)
		break
	}
	return result
}

/*
func main() {
	url := link{baseURL: "https://rasj.dk"}
	navBtn{name: "Forside", dir: "/", btnType: "button"}.HTML(&url)
	var mainNavBar navBar
	mainNavBar.appendBtn(navBtn{name: "Forside", dir: "/", btnType: "button"})
	mainNavBar.appendBtn(navBtn{name: "Help", dir: "/help", btnType: "button"})
	mainNavBar.appendBtn(navBtn{name: "Spil", dir: "/spil", btnType: "collapse"})
	mainNavBar.appendBtn(navBtn{name: "Wordle", dir: "/wordle", btnType: "button"})
	mainNavBar.appendBtn(navBtn{btnType: "end"})

	fmt.Println(mainNavBar.HTML())
}
*/
