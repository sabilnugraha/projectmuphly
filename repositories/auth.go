package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AddAdmin(admin models.Admin) (models.Admin, error)
	Login(username string) (models.Admin, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddAdmin(user models.Admin) (models.Admin, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(username string) (models.Admin, error) {
	var user models.Admin
	err := r.db.First(&user, "user_name=?", username).Error

	return user, err
}
