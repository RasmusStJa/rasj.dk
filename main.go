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

func getAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /about request")
	w.Header().Add("Content-Type", "text/html")

	var head header
	var pageBody, source, pinggy element
	source.CreateLink("here.", "github.com/RasmusStJa/rasj.dk")
	pinggy.CreateLink("", "pinggy.io")
	pageBody.CreateBody()

	pageBody.AppendChild(element{tag: "h1", innerText: "About"})

	pageBody.AppendChild(element{tag: "p", innerText: "This server is running on a Raspberry Pi 5 I have at home.<br>It's running Go for the backend, where I've written some code to generate the HTML. You can find that ",
		children: []element{source}})
	pageBody.AppendChild(element{tag: "p", innerText: "And because of my isp, I'm required to run this through a reverse proxy service, like ",
		children: []element{pinggy}})

	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func getPrime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /isnowaprime request")
	w.Header().Add("Content-Type", "text/html")

	getFactors := func(n int) []int {
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

	isPrime := func(n int) bool {
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

	var pageBody element
	pageBody.CreateBody()
	var head header

	now, _ := strconv.Atoi("20" + time.Now().Format("0601021504"))
	pageBody.AppendChild(element{tag: "h1", innerText: "Is now a prime?"})
	pageBody.AppendChild(element{tag: "p", innerText: "Is " + strconv.Itoa(now) + " a prime?"})

	if isPrime(now) {
		pageBody.AppendChild(element{tag: "p", innerText: "Yes"})
	} else {
		pageBody.AppendChild(element{tag: "p", innerText: "No"})
		factors := getFactors(now)
		if len(factors) > 0 {
			pageBody.AppendChild(element{tag: "p", innerText: "Here are its factors:"})
			list := element{tag: "ul"}

			for _, f := range factors {
				list.AppendChild(element{tag: "li", innerText: strconv.Itoa(f)})
			}
			pageBody.AppendChild(list)
		}
	}

	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func getFredag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /fredag request")

	var head header
	var pageBody element

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

	var head header
	var pageBody, navBar element

	pageBody.CreateBody()

	navBar.CreateNavBar()
	pageBody.AppendChild(element{tag: "h1", innerText: "Home page"})
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

	http.Handle("/file/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Got /file/ request for: %s\n", r.URL.Path)
		http.StripPrefix("/file/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	}))
	//serve any files in a static dir

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server is closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
