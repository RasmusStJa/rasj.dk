package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")
	var mainNavBar navBar
	mainNavBar.appendBtn(navBtn{name: "Forside", dir: "/", btnType: "button"})
	mainNavBar.appendBtn(navBtn{name: "Help", dir: "/help", btnType: "button"})
	mainNavBar.appendBtn(navBtn{name: "Spil", dir: "/spil", btnType: "collapse"})
	mainNavBar.appendBtn(navBtn{name: "Wordle", dir: "/wordle", btnType: "button"})
	mainNavBar.appendBtn(navBtn{btnType: "end"})

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, "<h1>Example page</h1>"+mainNavBar.HTML())
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server is closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
