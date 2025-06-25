package main

import (
	"com.int.recipies/http"
	"com.int.recipies/http/handlers"
)

func main() {
	recipieHandler := handlers.NewRecipieHandler(nil)
	handlers := http.Handlers{
		RecipieHandlers: *recipieHandler,
	}
	server := http.NewServer(handlers)
	server.ListenAndServe("localhost:8080")
}