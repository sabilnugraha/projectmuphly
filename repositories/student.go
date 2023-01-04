package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	AddStudent(product models.User) (models.User, error)
	GetStudent(Id int) (models.User, error)
}

func RepositoryStudent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddStudent(student models.User) (models.User, error) {
	err := r.db.Create(&student).Error

	return student, err
}

func (r *repository) GetStudent(Id int) (models.User, error) {
	var student models.User
	err := r.db.First(&student, Id).Error

	return student, err
}