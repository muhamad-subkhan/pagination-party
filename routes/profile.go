package routes

import (
	"party/handlers"
	"party/pkg/middleware"
	"party/pkg/mysql"
	"party/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profile := repositories.RepositoriesProfile(mysql.DB)
	h := handlers.HandlerProfile(profile)

	r.HandleFunc("/profile/{id}", middleware.Auth(middleware.UploadFile(h.GetProfile))).Methods("GET")
	r.HandleFunc("/profile-update/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("PATCH")
}