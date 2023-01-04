package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AddAdmin(admin models.Admin) (models.Admin, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddAdmin(user models.Admin) (models.Admin, error) {
	err := r.db.Create(&user).Error

	return user, err
}
