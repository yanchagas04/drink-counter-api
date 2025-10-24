package users

import (
	"drink-counter-api/users/routes"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Init(router *mux.Router, db *gorm.DB) {
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) { routes.CreateHandler(db, w, r)}).Methods("POST")
}