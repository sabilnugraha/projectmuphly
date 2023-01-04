package routes

import (
	"microservice/handlers"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryAuth(sql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.AddAdmin).Methods("POST")
}
