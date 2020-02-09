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
	handlers.Books = append(handlers.Books, models.Book{ID: "1", Title: "Война и Мир", Author: &models.Author{Firstname: "Лев", Lastname: "Толстой"}})
	handlers.Books = append(handlers.Books, models.Book{ID: "2", Title: "Преступление и наказание", Author: &models.Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
