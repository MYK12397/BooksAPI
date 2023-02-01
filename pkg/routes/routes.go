package routes

import (
	"github.com/MYK12397/BooksAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookRoutes = func(router *mux.Router) {

	router.HandlerFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookID}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/book/{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookID}", controllers.DeleteBook).Methods("DELETE")
}
