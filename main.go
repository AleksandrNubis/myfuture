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
	handlers.Books = append(handlers.Books, models.Book{ID: "1", Title: "Властелин Колец", Author: &models.Author{Firstname: "Джон", Lastname: "Толкиен"}})
	handlers.Books = append(handlers.Books, models.Book{ID: "2", Title: "Шерлок Холмс", Author: &models.Author{Firstname: "Артур", Lastname: "Дойл"}})
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
