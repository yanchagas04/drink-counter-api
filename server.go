package main

import (
	"drink-counter-api/driver"
	"drink-counter-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	utils.LoadEnv()
	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")
	ADDRESS := HOST + ":" + PORT
	db := driver.Connect()
	defer driver.Close(db)
	main_router := mux.NewRouter()
	main_router.HandleFunc("/", handler)
	log.Default().Println("Servidor rodando em " + ADDRESS)
	log.Fatal(http.ListenAndServe(":" + PORT, main_router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: "Hello World",
	})
}