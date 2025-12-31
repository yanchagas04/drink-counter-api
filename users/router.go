package users

import (
	"drink-counter-api/users/routes"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(router *mux.Router, db *gorm.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { routes.CreateHandler(db, w, r) }).Methods("POST")
	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.UpdateHandler(db, w, r) }).Methods("PUT")
	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.GetByIdHandler(db, w, r) }).Methods("GET")
	userRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { routes.GetHandler(db, w, r) }).Methods("GET")
	userRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { routes.LoginHandler(db, w, r) }).Methods("POST")
}
