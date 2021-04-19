package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
