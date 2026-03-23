package users

import (
	"drink-counter-api/users/routes"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(router *mux.Router, db *gorm.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) { routes.GetUsersHandler(db, w, r) }).Methods("GET") // get all users or by query
	userRouter.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) { routes.GetUserHandler(db, w, r) }).Methods("GET")
	userRouter.HandleFunc("/{username}", func(w http.ResponseWriter, r *http.Request) { routes.GetUserByUsername(db, w, r) }).Methods("GET") // get users by username
	userRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { routes.GetUserByUsername(db, w, r) }).Methods("GET") // representa a rota de cima, porém sem o parâmetro username, fazendo com que retorne um erro de falta de username

	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) { routes.CreateUserHandler(db, w, r) }).Methods("POST") // create user
	userRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { routes.UserLoginHandler(db, w, r) }).Methods("POST")

	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.UpdateUserHandler(db, w, r) }).Methods("PUT")

	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) { routes.DeleteUserHandler(db, w, r) }).Methods("DELETE")
}
