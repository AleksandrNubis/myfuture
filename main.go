package main

import (
	"github.com/gorilla/mux"
	"log"
	"myfuture/models"
	"net/http"
	"myfuture/handlers"
)

func main() {
	r := mux.NewRouter()
	models.books = append(books, models.Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	models.books = append(books, models.Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", handlers.getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.getBook).Methods("GET")
	r.HandleFunc("/books", handlers.createBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
