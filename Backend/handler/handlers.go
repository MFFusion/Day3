package handler

import (
	"encoding/json"
	"go-app/database"
	"go-app/model"
	"log"
	"net/http"
)

func SearchBooksHandler(w http.ResponseWriter, r *http.Request) {
	db := repository.Connect("bookstore.db")
	title := r.URL.Query().Get("title")
	log.Print(title)
	query := `SELECT Id,Title,Price,Author FROM books WHERE Title LIKE %` + "'title'" + `%`
	log.Print(query)
	rows := repository.GetRows(db, query)
	defer repository.Close(db)
	defer repository.CloseRows(rows)
	books := []model.Book{}
	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Price, &book.Author)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get all books
	db := repository.Connect("bookstore.db")
	rows := repository.GetRows(db, "SELECT Id,Title,Price,Author FROM books")
	defer repository.Close(db)
	defer repository.CloseRows(rows)
	books := []model.Book{}
	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Price, &book.Author)
		if err != nil {
			panic(err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}
