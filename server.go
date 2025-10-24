package main

import (
	"drink-counter-api/driver"
	"drink-counter-api/users"
	"drink-counter-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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
	driver.RunMigrations(db)
	defer driver.Close(db)
	main_router := mux.NewRouter()
	main_router.HandleFunc("/", handler).Methods("GET")
	users.Init(main_router, db)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	log.Default().Println("Servidor rodando em " + ADDRESS)
	log.Fatal(http.ListenAndServe(":" + PORT, cors(main_router)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{
		Message: "Hello World",
	})
}