package main

import (
	"database/sql"
	"fmt"
	_ "go-sqlite3"
	"html/template"
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, request *http.Request) {
	html, errorHome := template.ParseFiles("home.html")
	if errorHome != nil {
		fmt.Printf("ERROR: %x", errorHome)
	}
	html.Execute(writer, nil)
}

func user(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		html, errorUser := template.ParseFiles("user.html")
		if errorUser != nil {
			fmt.Printf("ERROR: %x", errorUser)
		}
		html.Execute(writer, nil)
	}
	if request.Method == "POST" {
		username := request.FormValue("username")
		password := request.FormValue("password")
		fmt.Println(username + password)
		db := connect()

		tx, txErr := db.Begin()
		if txErr != nil {
			log.Fatal(txErr)
		}

		cmd := "INSERT INTO users" +
			"(username, password)" +
			"VALUES" +
			"(?, ?)"

		stmt, stmtErr := tx.Prepare(cmd)
		if stmtErr != nil {
			log.Fatal(stmtErr)
		}
		stmt.Exec(username, password)
		tx.Commit()
		stmt.Close()
		db.Close()
		http.Redirect(writer, request, "/login", http.StatusFound)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	html, errorUser := template.ParseFiles("login.html")
	if errorUser != nil {
		fmt.Printf("ERROR: %x", errorUser)
	}
	html.Execute(writer, nil)
}

func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./Cumro.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("localhost:5000", nil)
	log.Fatal(err)
}
