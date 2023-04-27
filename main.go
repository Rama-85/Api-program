package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Isbn          string  `json:"isbn"`
	Discount      int     `json:"discount"`
	Typeofthebook string  `json:"typeofthebook"`
	Price         *Price  `json:"price"`
	Author        *Author `json:"author"`
}

type Price struct {
	Onlineprice  string `json:"onlineprice"`
	Offlineprice string `json:"offlineprice"`
}

type Author struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Noofbookswritten int    `json:"noofbookswritten"`
	Place            string `json:"place"`
	Language         string `json:"language"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.Id == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode("Item not found")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{Id: 1, Name: "Golang", Isbn: "yes", Discount: 10, Typeofthebook: "Medium", Author: &Author{Id: 10, Name: "chand", Noofbookswritten: "100", Place: "Hyderabad", Language: "Telugu"}, Price: &Price{Onlineprice: "1000", Offlineprice: "800"}})
	books = append(books, Book{Id: 2, Name: "Github", Isbn: "yes", Discount: 40, Typeofthebook: "Large", Author: &Author{Id: 20, Name: "Dennis", Noofbookswritten: "500", Place: "London", Language: "English"}, Price: &Price{Onlineprice: "900", Offlineprice: "800"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{ID}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
