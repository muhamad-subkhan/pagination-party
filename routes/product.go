package routes

import (
	"party/handlers"
	"party/pkg/middleware"
	"party/pkg/mysql"
	"party/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router){
	productRepositories := repositories.RepositoriesProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepositories)

	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product",middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product-update/{id}",middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
	r.HandleFunc("/product", h.FindProduct).Methods("GET")
}