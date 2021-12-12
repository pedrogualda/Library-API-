package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Book Struct (Model)
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Get all books
//Every function that you create that is arout handler has to have
//These two parameters: Request and Response.
func getBooks(w http.ResponseWriter, r *http.Request)  {
	//Now we wanna set the header value of content type to application/json
	//Or else it's just gonna get served up as text
	w.Header().Set("Content type", "application/json")
	json.NewEncoder(w).Encode(books)
}
//Get a single book
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //This will get any parameters
	//Then we need to loop through the books and find the correct id
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
		json.NewEncoder(w).Encode(&Book{})
	}

}
//Create a new book
func createBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range books{
			if item.ID == params["id"] {
				books = append(books[:index], books[index+1:]...)
				w.Header().Set("Content type", "application/json")
				var book Book
				_ = json.NewDecoder(r.Body).Decode(&book)
				book.ID =  params["id"]
				books = append(books, book)
				json.NewEncoder(w).Encode(book)
				return
			}
		}
		json.NewEncoder(w).Encode(books)
}

//Delete book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//author Struct
type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Init books var as a slice of Book Struct
//Slice if basically a variable length array. A slice is variable in the
//Amount of values it contains

var books []Book

func main()  {
	// Init Router
	rout := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &Author {Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "971234", Title: "Book two", Author: &Author {Firstname: "Medro", Lastname: "Potta"}})

	//Create Route Handlers, which establishes
	//Endpoints for our APIs
	//We use .methods to show which type of HTTP method to use.
	rout.HandleFunc("/api/books", getBooks).Methods("GET")
	rout.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	rout.HandleFunc("/api/books", createBook).Methods("POST")
	rout.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	rout.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", rout))
	//Incluindo a função ListenAndServe dentro da função log.Fatal, pois caso ela não funcione, retorna erro.

}