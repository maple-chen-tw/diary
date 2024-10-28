package main

import (
	"diary/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.Router()
	log.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
