package routes

import (
	"party/handlers"
	"party/pkg/mysql"
	"party/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	Auth := repositories.RepositoriesAuth(mysql.DB)
	h := handlers.HandlerAuth(Auth)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
}