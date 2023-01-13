package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	JournalRoutes(r)
	StudentRoutes(r)
	BillingRoutes(r)
	NotificationRoutes(r)
}
