package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type BillingRepository interface {
	AddBilling(product models.BillingMonthly) (models.BillingMonthly, error)
}

func RepositoryBilling(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddBilling(billing models.BillingMonthly) (models.BillingMonthly, error) {
	err := r.db.Create(&billing).Error

	return billing, err
}
