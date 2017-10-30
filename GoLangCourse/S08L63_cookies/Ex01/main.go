package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)

	http.ListenAndServe(":8080", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	fmt.Println("\nEntering the cookie jar")
	c, err := req.Cookie("access")
	if err != nil {
		fmt.Println("We have an error", err)
		log.Println(err)
	}

	var i int
	if c == nil {
		fmt.Println("Where's the cookie?")
		i = 0
	} else {
		fmt.Println(c.String())

		i, err = strconv.Atoi(c.Value)
		if err != nil {
			fmt.Println("Cannot convert cookie to number", err)
			i = 0
		}
	}

	http.SetCookie(res, &http.Cookie{
		Name:   "access",
		Value:  strconv.Itoa(i + 1),
		MaxAge: 3600,
	})

	fmt.Fprintln(res, "You were here before", i, "times")

}
