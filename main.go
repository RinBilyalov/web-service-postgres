package main

import (
	"github.com/joho/godotenv"
	"net/http"
	"web-service-postgres/config"
)

func main() {

	godotenv.Load()

	server := &http.Server{
		Addr: ":8080",
	}

	config.ConnectDatabase()
	server.ListenAndServe()
}
