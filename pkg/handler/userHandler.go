package handler

import (
	"net/http"
	"web-service-postgres/cmd"

	"github.com/gorilla/mux"
)

func NewUserHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users", cmd.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", cmd.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}/comments", cmd.GetUserComments).Methods("GET")
	router.HandleFunc("/user/{id}/comments/{id}", cmd.GetUserCommentById).Methods("GET")

	//router.HandleFunc("/user/{id}", cmd.CreateUser).Methods("POST")

	return router
}
