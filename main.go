package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	books = append(books, Books{ID: "1", Isbn: "mahendra", Title: "MSR", Author: &Author{FirstName: "msr", LastName: "MSR"}})
	books = append(books, Books{ID: "2", Isbn: "mahendra singh", Title: "MSR MSR", Author: &Author{FirstName: "msr", LastName: "MSR"}})

	r := mux.NewRouter()

	r.HandleFunc("/v1/api/books", getBooks).Methods("GET")
	r.HandleFunc("/v1/api/books/{id}", getBooksByID).Methods("GET")
	r.HandleFunc("/v1/api/books/createbooks", CreateBooks).Methods("POST")
	r.HandleFunc("/v1/api/books/BookUpdate/{id}", BookUpdate).Methods("PUT")
	r.HandleFunc("/v1/api/books/delete/{id}", DeletBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
