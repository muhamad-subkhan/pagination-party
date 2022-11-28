package routes

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	ProfileRoutes(r)
	AuthRoutes(r)
	ProductRoutes(r)
}