package controllers

import (
	"BookStoreAPIs/core/models"
	"BookStoreAPIs/core/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["book_id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var newBook models.Book
	json.Unmarshal(body, &newBook)
	//utils.ParseBody(r, newBook)
	fmt.Println("After Parsing: ", newBook)
	book := newBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var bookPayload = &models.Book{}
	utils.ParseBody(r, bookPayload)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["book_id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	oldBook, db := models.GetBookById(id)
	oldBook.Name = bookPayload.Name
	oldBook.Author = bookPayload.Author
	oldBook.Publication = bookPayload.Publication
	db.Save(oldBook)
	res, _ := json.Marshal(oldBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["book_id"], 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
