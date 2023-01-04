package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type JournalRepository interface {
	ReadJournals() ([]models.JournalView, error)
	InputJournal(product models.Journal) (models.Journal, error)
	GetJournal(Id int) (models.Journal, error)
}

func RepositoryJournal(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ReadJournals() ([]models.JournalView, error) {
	var journal []models.JournalView
	err := r.db.Raw(`SELECT Id, id_transaction, date, ops_input, ops_output, monthly_input, monthly_output, mahad_input, mahad_output, uraian, (ops_input + mahad_input + monthly_input) AS debet, (ops_output + mahad_output + monthly_output) AS Kredit,
	SUM((ops_input + mahad_input + monthly_input - ops_output - mahad_output - monthly_output)) OVER (ORDER BY date) AS saldoakhir
	FROM journals`).Scan(&journal).Error

	return journal, err
}

func (r *repository) InputJournal(journal models.Journal) (models.Journal, error) {
	err := r.db.Create(&journal).Error

	return journal, err
}

func (r *repository) GetJournal(Id int) (models.Journal, error) {
	var journal models.Journal
	err := r.db.First(&journal, Id).Error

	return journal, err
}
