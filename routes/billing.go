package routes

import (
	"microservice/handlers"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func BillingRoutes(r *mux.Router) {

	billingRepository := repositories.RepositoryBilling(sql.DB)

	h := handlers.HandlerBilling(billingRepository)

	r.HandleFunc("/addbilling", h.AddBilling).Methods("POST")

}
