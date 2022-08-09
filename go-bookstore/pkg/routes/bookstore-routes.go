package routes

import (
	"github.com/gorilla/mux"
	"github.com/rosebash/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("POST")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
