package routes

import (
	"microservice/handlers"
	"microservice/pkg/middleware"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {

	userRepository := repositories.RepositoryAuth(sql.DB)
	h := handlers.HandlerAuth(userRepository)

	r.HandleFunc("/register", h.AddAdmin).Methods("POST")
	r.HandleFunc("/studentregister/{id}", h.CreateStudentAccount).Methods("PATCH")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/lastnis", h.FindLastNis).Methods("GET")
	r.HandleFunc("/studentlogin", h.LoginStudent).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
	r.HandleFunc("/check-auth-stud", middleware.Auth(h.CheckAuthStudentId)).Methods("GET")
}
