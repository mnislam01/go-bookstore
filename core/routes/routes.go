package routes

import (
	"BookStoreAPIs/core/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

var RoutesRegistry = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{book_id}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{book_id}", controllers.UpdateBookById).Methods(http.MethodPut)
	router.HandleFunc("/books/{book_id}", controllers.DeleteBookById).Methods(http.MethodDelete)
}
