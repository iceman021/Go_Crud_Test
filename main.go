package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"time/date"
	"strconv"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Book struct {
	ID              string    `json:"id" validate:"alpha"`
	Title           string    `json:"title" validate:"alpha"`
	Description     string    `json:"description" validate:"alpha"`
	PublicationDate time.Time `json:"publicationdate"`
}

type Author struct {
	ID      string    `json:"id"`
	Name    string `json:"name" validate:"alpha"`
	Surname string `json:"surname" validate:"alpha"`
}

//conection between authors and books
type BookAuthor struct {
	Author        Author `json:"author" validate:"alpha"`
	PublishedBook Book   `json:"publishedbook" validate:"alpha"`
}

//type Library struct {
//	BookAuthors []BookAuthor `json:"bookauthors"`
//}

// TODO: When you delete author or book you will have to delete from bookAuthor + on deleting author, delete from bookAuthor + all books

var books []Book
var authors []Author

 func getBooks(w http.ResponseWriter, router *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	json.NewEncoder(w).Encode(books)
 }

// // Get A Single Book
 func getBook(w http.ResponseWriter, router *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	params := mux.Vars(router) // Get Params
 	//Loop through Books and find with Id
 	for _, item := range books {
 		if item.ID == params["id"] {
 			json.NewEncoder(w).Encode(item)
 			return
 		}
 	}
 }

// Create a New Book
 func createBook(w http.ResponseWriter, router *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	var book Book
 	_ = json.NewDecoder(router.Body).Decode(&book)
 	book.ID = "BookOne" // MOCK ID - NOT SAFE IN PRODUCTION
 	books = append(books, book)
 	json.NewEncoder(w).Encode(book)
 }

 func updateBook(w http.ResponseWriter, router *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	params := mux.Vars(router) // Get Params
 	//Loop through Books and find with Id
 	for index, item := range books {
 		if item.ID == params["id"] {
 			books = append(books[:index], books[index+1:]...)
 			var book Book
 			_ = json.NewDecoder(router.Body).Decode(&book)
 			book.ID = params["id"]
 			books = append(books, book)
 			json.NewEncoder(w).Encode(book)
 			return
 		}
 	}
 	json.NewEncoder(w).Encode(books)
 }

 func deleteBook(w http.ResponseWriter, router *http.Request) {
 	w.Header().Set("Content-Type", "application/json")
 	params := mux.Vars(router) // Get Params
 	//Loop through Books and find with Id
 	for index, item := range books {
 		if item.ID == params["id"] {
 			books = append(books[:index], books[index+1:]...)
 			break
 		}
 	}
 	json.NewEncoder(w).Encode(books)
 }

var validate *validator.Validate

func main() {

	 var authors []string
	 authors = append(authors, "Mark Twain", "Charles Dickens", "Franz Kafka", "Agatha Christy")

	var publishedbooks []Book
	authors = append(authors, Author{ID: 5, Name: "Mark", Surname: "Twain", PublishedBooks: publishedbooks})

	books = append(books, Book{ID: "4", Title: "Book One", Description: "345543", Authors: authors, PublicationDate: time.Now()})

	for _, item := range BookAuthor {
		item.PublishedBook = books
		authors = append(authors, item)
	}

	fmt.Print(authors)

	 books = append(books, Book{ID: "2", Title: "Book Two", Description: "jyjty",
	 	Author: &Author{Name: "Franz", Surname: "Kafka"}, PublicationDate: time.Date()})

	 books = append(books, Book{ID: "3", Title: "Book One", Description: "ujyt",
	 	Author: &Author{Name: "Charles", Surname: "Dickens"}, PublicationDate: time.Date()})

	 books = append(books, Book{ID: "4", Title: "Book Two", Description: "kuu",
	 	//Author: &Author{Name: "Agatha", Surname: "Christy"}, PublicationDate: time.date()})

	r := mux.NewRouter()

	 r.HandleFunc("/books", getBooks).Methods("GET")
	 r.HandleFunc("/bookss/id{id}", getBook).Methods("GET")
	 r.HandleFunc("/books", createBook).Methods("POST")
	 r.HandleFunc("/bookss/{id}", updateBook).Methods("PUT")
	 r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("starting server at PORT 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

/*
package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/akhil/go-bookstore/pkg/utils"
	"github.com/akhil/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//get book
func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//create
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//delete
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//update
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
*/


