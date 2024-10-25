package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"web-service-postgres/config"
	"web-service-postgres/pkg/handler"
)

func main() {

	godotenv.Load()

	handlerUser := handler.NewUserHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlerUser,
	}

	config.ConnectDatabase()
	server.ListenAndServe()
}
