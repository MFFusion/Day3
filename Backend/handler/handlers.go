package handler

import (
	"encoding/json"
	"go-app/database"
	"net/http"
)

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get all books
	db := repository.Connect("books.db")
	rows := repository.GetRows(db, "SELECT * FROM books")
	defer repository.Close(db)
	defer repository.CloseRows(rows)
	books := []Book{}
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}
