package routes

import (
	"microservice/handlers"
	"microservice/pkg/middleware"
	"microservice/pkg/sql"

	"microservice/repositories"

	"github.com/gorilla/mux"
)

func NotificationRoutes(r *mux.Router) {

	notificationRepository := repositories.RepositoryNotification(sql.DB)

	h := handlers.HandlerNotification(notificationRepository)

	r.HandleFunc("/addnotification", h.AddNotification).Methods("POST")
	r.HandleFunc("/notification", h.GetNotification).Methods("GET")
	r.HandleFunc("/position", h.GetAllAdminByPosition).Methods("POST")
	r.HandleFunc("/adminnotification", middleware.Auth(h.FindNotificationtId)).Methods("GET")

}
