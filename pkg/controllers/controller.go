package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cleave3/bookstore/pkg/models"
	"github.com/cleave3/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var book models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {

	books := models.GetAllBooks()

	utils.HandleSucess(w, books)
}


func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	bookId := vars["bookId"];

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, _ := models.GetBookById(id)

	utils.HandleSucess(w, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}

	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()

	utils.HandleCreated(w, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	bookId := vars["bookId"];

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(id)

	utils.HandleSucess(w, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}

	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)

	bookId := vars["bookId"];

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, db := models.GetBookById(id)

	if updateBook.Name != "" {
		
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		
		book.Publication = updateBook.Publication
	}


	db.Save(&book)

	utils.HandleSucess(w, book)
}