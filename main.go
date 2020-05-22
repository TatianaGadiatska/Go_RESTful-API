package main

import (
	"./controllers"
	"./driver"
	"./models"
	"database/sql"
	_ "database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/subosito/gotenv"
	"log"
	"net/http"
)

var books []models.Book
var db *sql.DB

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
