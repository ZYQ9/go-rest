package main

import (
	"net/http"

	"github.com/ZYQ9/go-rest/controllers"
	"github.com/ZYQ9/go-rest/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	models.ConnectDatabase()

	server.ListenAndServe()
}
