package routes

import (
	"github.com/abhikrishnaram/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookbyId).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
