package routes

import (
	"microservice/handlers"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func StudentRoutes(r *mux.Router) {

	studentRepository := repositories.RepositoryStudent(sql.DB)
	h := handlers.HandlerStudent(studentRepository)

	r.HandleFunc("/addstudent", h.AddStudent).Methods("POST")
}
