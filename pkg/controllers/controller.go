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

	utils.HandleSucess(w, 200, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, _ := models.GetBookById(id)

	if book.ID != uint(id) {
		utils.HandleBadRequest(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.HandleSucess(w, 200, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	newBook := &models.Book{}

	utils.ParseBody(r, newBook)

	b := newBook.CreateBook()

	utils.HandleSucess(w, 201, b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book := models.DeleteBook(id)

	utils.HandleSucess(w, 200, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}

	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	book, db := models.GetBookById(id)

	if book.ID != uint(id) {
		utils.HandleBadRequest(w, http.StatusNotFound, "Book not found")
		return
	}

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

	utils.HandleSucess(w, 200, book)
}
