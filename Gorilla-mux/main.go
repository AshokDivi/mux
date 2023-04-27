package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Booktype string  `json:"booktype"`
	Price    *Price  `json:"price"`
	Author   *Author `json:"director"`
}

type Author struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Bookswritten int    `json:"bookswritten"`
	Place        string `json:"place"`
	Language     string `json:"language"`
}

type Price struct {
	Onlineprice  int `json:"onlineprice"`
	Offlineprice int `json:"Offlineprice"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["ID"] {
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
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

var books []Book

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", Name: "Think like a monk", Booktype: "Peace", Author: &Author{ID: "1", Name: "Jay Setty", Bookswritten: 5, Place: "United states", Language: "English"}, Price: &Price{Onlineprice: 300, Offlineprice: 500}})
	books = append(books, Book{ID: "2", Name: "Rich dad Poor dad", Booktype: "Peace", Author: &Author{ID: "2", Name: "Robert kiyosaki", Bookswritten: 14, Place: "United States", Language: "English"}, Price: &Price{Onlineprice: 450, Offlineprice: 600}})
	books = append(books, Book{ID: "3", Name: "Adhinetha", Booktype: "Peace", Author: &Author{ID: "3", Name: "Pattabhiram", Bookswritten: 35, Place: "India", Language: "English"}, Price: &Price{Onlineprice: 350, Offlineprice: 300}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{ID}", getBook).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

	fmt.Println("Starting server on 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
