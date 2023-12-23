package model

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Price  string `json:"price"`
	Author string `json:"author"`
}
