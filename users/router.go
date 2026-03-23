package users

import (
	"drink-counter-api/users/routes"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(router *mux.Router, db *gorm.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) { routes.GetHandler(db, w, r) }).Methods("GET")
	userRouter.HandleFunc("/{username}", func(w http.ResponseWriter, r *http.Request) { routes.GetByUsername(db, w, r) }).Methods("GET")
	userRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { routes.GetByUsername(db, w, r) }).Methods("GET") // representa a rota de cima, porém sem o parâmetro username, fazendo com que retorne um erro de falta de username

	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) { routes.CreateHandler(db, w, r) }).Methods("POST")
	userRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { routes.LoginHandler(db, w, r) }).Methods("POST")

	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.UpdateHandler(db, w, r) }).Methods("PUT")

	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.DeleteHandler(db, w, r) }).Methods("DELETE")
}
