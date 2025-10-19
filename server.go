package main

import (
	"drink-counter-api/driver"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := driver.Connect()
	log.Default().Println("Server is running on http://localhost:8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer driver.Close(db)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}