package main

import (
	"fmt"
	//"log"
	//"net/http"
)

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
