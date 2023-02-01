package routes

import (
	"microservice/handlers"
	"microservice/pkg/middleware"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func StudentRoutes(r *mux.Router) {

	studentRepository := repositories.RepositoryStudent(sql.DB)
	h := handlers.HandlerStudent(studentRepository)

	r.HandleFunc("/addstudent", h.AddStudent).Methods("POST")
	r.HandleFunc("/addclass", h.AddGroupClass).Methods("POST")
	r.HandleFunc("/addstudenttoclass", h.AddStudentToGroupClass).Methods("POST")
	r.HandleFunc("/addsubclass", h.AddSubClass).Methods("POST")
	r.HandleFunc("/addstudenttosubclass", h.AddStudentToSubClass).Methods("POST")
	r.HandleFunc("/addphoto/{id}", middleware.UploadFile(h.AddPhoto)).Methods("PATCH")
	r.HandleFunc("/student/{nis}", h.GetNIS).Methods("GET")
}
