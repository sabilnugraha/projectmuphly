package routes

import (
	"microservice/handlers"
	"microservice/pkg/sql"

	"microservice/repositories"

	"github.com/gorilla/mux"
)

func NotificationRoutes(r *mux.Router) {

	notificationRepository := repositories.RepositoryNotification(sql.DB)

	h := handlers.HandlerNotification(notificationRepository)

	r.HandleFunc("/addnotification", h.AddNotification).Methods("POST")
}
