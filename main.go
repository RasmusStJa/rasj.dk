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

const stylepth string = "style.css"
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
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")
	head := header{stylepath: stylepth}

	pageBody := element{tag: "body"}

	pageBody.AppendChild(element{tag: "h1", innerText: "Example page"})
	pageBody.AppendChild(element{tag: "p", innerText: "The time is currently " + time.Now().Format("15:04")})
	pageBody.AppendChild(element{tag: "button", innerText: "Is now a prime?", attributes: []attribute{{name: "onclick", value: "location.href=\"https://pi.rasj.dk/isnowaprime\""}}})
	
	w.Header().Add("Content-Type", "text/html")
	io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
	//io.WriteString(w, ().HTML())
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
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

	w.Header().Add("Content-Type", "text/html")

	pageBody := element{tag: "body"}
	head := header{stylepath: stylepth}

	now, _ := strconv.Atoi(time.Now().Format("0601021504"))
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
		var primes []element
		offset := 1
		for len(primes) < 3 {
			d1 := time.Now().Add(time.Minute * time.Duration(offset))
			d2, _ := strconv.Atoi(d1.Format("0601021504"))
			if isprime(d2) {
				primes = append(primes, element{tag: "p", innerText: d1.Format("2006-01-02 15:04")})
			}
			offset++
		}
		pageBody.AppendChild(element{tag: "p", innerText: "Here are the next 3 prime times:"})
		pageBody.AppendChildren(primes)*/
 io.WriteString(w, head.HTML())
	io.WriteString(w, pageBody.HTML())
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/isnowaprime", getPrime)

	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server is closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
