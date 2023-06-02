package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var books []Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBooksByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	for _, item := range books {
		if item.ID == params {

			json.NewEncoder(w).Encode(item)
			return

		}
	}

	json.NewEncoder(w).Encode(&Books{})

}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Books

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func BookUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	for index, item := range books {

		if item.ID == params {
			books = append(books[:index], books[index+1:]...)
			var book Books

			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}

	json.NewEncoder(w).Encode(books)

}

func DeletBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]

	for index, item := range books {

		if item.ID == params {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}
