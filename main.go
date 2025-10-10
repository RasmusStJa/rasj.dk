package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
	func getNavbar() navBar {
		var mainNavBar navBar
		mainNavBar.AppendBtn(navBtn{name: "Forside", dir: "/", btnType: "button"})
		mainNavBar.AppendBtn(navBtn{name: "Help", dir: "/help", btnType: "button"})
		mainNavBar.AppendBtn(navBtn{name: "Spil", dir: "/spil", btnType: "collapse"})
		mainNavBar.AppendBtn(navBtn{name: "Wordle", dir: "/wordle", btnType: "button"})
		mainNavBar.AppendBtn(navBtn{btnType: "end"})
		return mainNavBar
	}
*/
func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /about request")
	w.Header().Add("Content-Type", "text/html")

	head := header{}
	pageBody := element{}
	pageBody.CreateBody()

	pageBody.AppendChild(element{tag: "h1", innerText: "About"})
	pageBody.AppendChild(element{tag: "p", innerText: "This server is running on a Raspberry Pi 5 I have at home.<br>It's running Go for the backend, where I've written some code to generate the HTML. You can find that ",
		children: []element{{tag: "a", innerText: "here.", attributes: []attribute{{name: "href", value: "https://pi.rasj.dk/source"}}}}})

	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func isprime(n int) bool {
	if n < 2 || n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getFactors(n int) []int {
	if n < 2 || n%2 == 0 {
		return []int{}
	}
	if n == 2 {
		return []int{2}
	}
	factors := []int{}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	if n > 1 {
		factors = append(factors, n/factors[len(factors)-1])
	}
	return factors
}

func getPrime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /isnowaprime request")

	pageBody := element{}
	pageBody.CreateBody()
	head := header{}

	now, _ := strconv.Atoi("20" + time.Now().Format("0601021504"))
	pageBody.AppendChild(element{tag: "h1", innerText: "Is now a prime?"})
	pageBody.AppendChild(element{tag: "p", innerText: "Is " + strconv.Itoa(now) + " a prime?"})
	if isprime(now) {
		pageBody.AppendChild(element{tag: "p", innerText: "Yes"})
	} else {
		pageBody.AppendChild(element{tag: "p", innerText: "No"})
		factors := getFactors(now)
		if len(factors) > 0 {
			pageBody.AppendChild(element{tag: "p", innerText: "Here are its factors:"})
			list := element{tag: "ul", children: []element{}}

			for _, f := range factors {
				list.AppendChild(element{tag: "li", innerText: strconv.Itoa(f)})
			}
			pageBody.AppendChild(list)
		}
	}
	/*
		var prime int
		var offset uint16 = 1
		for {
			d1 := time.Now().Add(time.Minute * time.Duration(offset))
			d2, _ := strconv.Atoi(d1.Format("0601021504"))
			if isprime(d2) {
				prime = d2
				break
			}
			offset++
		}
		pageBody.AppendChild(element{tag: "p", innerText: "Here is the next prime:"})
		pageBody.AppendChild(element{tag: "p", innerText: strconv.Itoa(prime)})*/
	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func getFredag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /fredag request")

	head := header{}
	pageBody := element{}

	pageBody.CreateBody()

	pageBody.AppendChild(element{tag: "h1", innerText: "Er det fredag idag?"})
	//day,_ := strconv.Atoi(time.Weekday())
	p := element{tag: "p"}
	if int(time.Now().Weekday()) == 5 {
		p.innerText = "JAAAA!!!!!!!!"
	} else {
		p.innerText = "nej >:("
	}
	pageBody.AppendChild(p)

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")

	head := header{}
	pageBody := element{}
	navBar := element{}

	pageBody.CreateBody()

	navBar.CreateNavBar()
	pageBody.AppendChild(element{tag: "h1", innerText: "Example page"})
	//pageBody.AppendChild(about)
	//pageBody.AppendChild(nowaprime)
	pageBody.AppendChild(navBar)
	pageBody.AppendChild(element{tag: "p", innerText: "The time is currently " + time.Now().Format("15:04")})

	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/about", getAbout)
	http.HandleFunc("/fredag", getFredag)
	http.HandleFunc("/isnowaprime", getPrime)
	http.HandleFunc("/source", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got /source request")
		http.Redirect(w, r, "https://github.com/RasmusStJa/rasj.dk", http.StatusMovedPermanently)
	})

	http.Handle("/", http.FileServer(http.Dir("./static"))) //serve any files in a "static" dir

	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server is closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
