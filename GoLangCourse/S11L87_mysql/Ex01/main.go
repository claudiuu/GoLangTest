package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "")
	lg(err)
	defer db.Close()

	err = db.Ping()
	lg(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	lg(err)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(res, "Database works!")
	lg(err)
}

func lg(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
