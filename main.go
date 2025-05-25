package main

import (
	"book-management-api/book"
	"log"
	"net/http"
)

func main() {
	store := book.NewStore()
	book.StartLogger()
	handler := book.MakeHandler(store)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}