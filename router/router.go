package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/user", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.EditUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/user/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
