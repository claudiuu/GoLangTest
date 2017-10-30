package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template
var dbUser map[string]*user
var dbSession map[string]string

type user struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

func init() {
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	dbUser = map[string]*user{}
	dbUser["claudiu"] = &user{
		"Claudiu",
		"Mateias",
		"claudiu",
		"claudiu@claudiuu.com",
		"abc",
	}
	dbSession = map[string]string{}

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/account", account)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	fmt.Println("index")

	usr := getUserFromCookie(res, req)

	fmt.Println("Identified user", usr)
	templates.ExecuteTemplate(res, "index.gohtml", usr)
}

func signup(res http.ResponseWriter, req *http.Request) {
	fmt.Println("signup")
	var usr *user
	if req.Method == http.MethodPost {
		fmt.Println("received Post method")
		fName := req.FormValue("fname")
		lName := req.FormValue("lname")
		email := req.FormValue("email")
		username := req.FormValue("username")

		usr = &user{
			fName,
			lName,
			username,
			email,
			"",
		}
		dbUser[username] = usr
		sessionID := "session_" + username
		dbSession[sessionID] = username
		http.SetCookie(res, &http.Cookie{
			Name:  "session",
			Value: sessionID,
		})
		fmt.Println("Saved cookie session")
		http.Redirect(res, req, "/", http.StatusSeeOther)
	} else {
		usr = getUserFromCookie(res, req)
	}
	templates.ExecuteTemplate(res, "signup.gohtml", usr)
}

func login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("login")
	var usr *user
	if req.Method == http.MethodPost {
		fmt.Println("received Post method")
		newUsername := req.FormValue("username")
		sessionID := "session_" + newUsername
		dbSession[sessionID] = newUsername
		http.SetCookie(res, &http.Cookie{
			Name:  "session",
			Value: sessionID,
		})
		fmt.Println("Saved cookie session")
		usr = dbUser[newUsername]

		http.Redirect(res, req, "/", http.StatusSeeOther)
	} else {
		usr = getUserFromCookie(res, req)
	}

	templates.ExecuteTemplate(res, "login.gohtml", usr)
}

func account(res http.ResponseWriter, req *http.Request) {
	fmt.Println("account")
	usr := getUserFromCookie(res, req)
	templates.ExecuteTemplate(res, "account.gohtml", usr)
}

func logout(res http.ResponseWriter, req *http.Request) {
	fmt.Println("logout")
	usr := getUserFromCookie(res, req)
	if usr != nil {
		c, _ := req.Cookie("session")
		delete(dbSession, c.Value)
		http.SetCookie(res, &http.Cookie{
			Name:   "session",
			Value:  "",
			MaxAge: -1,
		})
	}
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func getUserFromCookie(res http.ResponseWriter, req *http.Request) (usr *user) {
	c, err := req.Cookie("session")
	if err != nil {
		// http.Redirect(res, req, "/", http.StatusSeeOther)
		fmt.Println("No cookie named session")
		return
	}
	username, ok := dbSession[c.Value]

	if !ok {
		fmt.Println("No username found for session", c.Value)
		// http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	usr, ok = dbUser[username]
	if !ok {
		fmt.Println("No user found for username", username)
		// http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("cookie", c.Value, "username", username, "user", usr)
	return usr
}
