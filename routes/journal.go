package routes

import (
	"microservice/handlers"
	"microservice/pkg/sql"
	"microservice/repositories"

	"github.com/gorilla/mux"
)

func JournalRoutes(r *mux.Router) {

	journalRepository := repositories.RepositoryJournal(sql.DB)

	h := handlers.HandlerJournal(journalRepository)

	r.HandleFunc("/journal", h.InputJournal).Methods("POST")
	r.HandleFunc("/journals", h.ReadJournals).Methods("GET")
}
