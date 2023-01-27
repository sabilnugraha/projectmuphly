package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	AddStudent(user models.User) (models.User, error)
	GetStudent(Id int) (models.User, error)
	AddPhoto(product models.User) (models.User, error)
	GetNIS(Nis int) (models.User, error)
	AddGroupClass(class models.Groupclass) (models.Groupclass, error)
	AddSubClass(class models.SubClass) (models.SubClass, error)
}

func RepositoryStudent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddStudent(student models.User) (models.User, error) {
	err := r.db.Create(&student).Error

	return student, err
}

func (r *repository) AddGroupClass(class models.Groupclass) (models.Groupclass, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) AddSubClass(class models.SubClass) (models.SubClass, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) GetStudent(Id int) (models.User, error) {
	var student models.User
	err := r.db.First(&student, Id).Error

	return student, err
}

func (r *repository) GetNIS(Nis int) (models.User, error) {
	var student models.User
	err := r.db.First(&student, "nis = ?", Nis).Error

	return student, err
}

func (r *repository) AddPhoto(student models.User) (models.User, error) {
	err := r.db.Save(&student).Error

	return student, err
}
