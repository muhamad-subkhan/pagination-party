package main

import (
	"fmt"
	"net/http"
	"party/database"
	"party/pkg/mysql"
	"party/routes"

	"github.com/gorilla/mux"
)

func main() {

	mysql.Database()
	
	database.Migration()

	r := mux.NewRouter()

	routes.Routes(r.PathPrefix("/api").Subrouter())
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Server Running on Port: 3000")

	http.ListenAndServe("localhost:3000", r)
}