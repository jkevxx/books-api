package main

import (
	"books-api/internal/service"
	"books-api/internal/store"
	"books-api/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Connect to SQLLite

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exist
	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	// Inject our dependencies
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	// Configure paths
	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Server listening in http://localhost:8080")

	// Listening the server

	log.Fatal(http.ListenAndServe(":8080", nil))

}
