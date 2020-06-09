/**
	https://www.youtube.com/watch?v=SonwZ6MF5BE
	Simple REST API to get, add, update and delete books with (id, isbn, title, author {firstname, lastname}
	GET, POST, PUT and DELETE

	To Run : "go build && ./RestApi"
	Postman : GET - http://localhost:8000/api/books
	Postman : GET - http://localhost:8000/api/books/1

*/

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books []Book

type Book struct {
	ID string 	`json:"id"`
	Isbn string	`json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string 	`json:"firstname"`
	Lastname string		`json:"lastname"`
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)  //Get params
	//Loop through books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	//_ = json.NewEncoder(writer).Encode(&Book{})
}

func createBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	_ = json.NewEncoder(writer).Encode(book)
}

func updateBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
		}
	}
	json.NewEncoder(writer).Encode(books)
}

func deleteBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(writer).Encode(books)
	}
}

func main() {
	//Init router
	r := mux.NewRouter()

	//Mock data
	books = append(books, Book{ID: "1", Isbn: "2423423", Title: "Book One", Author: &Author{Firstname: "John", Lastname:"Doe"}})
	books = append(books, Book{ID: "2", Isbn: "2554566", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname:"Smith"}})


	//Route Handler / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}







