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
	AddStudentToGroupClass(class models.ClassUser) (models.ClassUser, error)
	AddSubClass(class models.SubClass) (models.SubClass, error)
	AddStudentToSubClass(class models.SubclassUser) (models.SubclassUser, error)
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

func (r *repository) AddStudentToGroupClass(class models.ClassUser) (models.ClassUser, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) AddSubClass(class models.SubClass) (models.SubClass, error) {
	err := r.db.Create(&class).Error

	return class, err
}

func (r *repository) AddStudentToSubClass(class models.SubclassUser) (models.SubclassUser, error) {
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

func GetUserClassInfoByUserID(userID uint, db *gorm.DB) ([]models.UserClassDetails, error) {
	var userClassInfos []models.UserClassDetails
	err := db.Table("class_users").
		Select("users.full_name, users.nik, users.nis, groupclasses.startclass, groupclasses.endclass").
		Joins("JOIN users ON class_users.user_id = users.id").
		Joins("JOIN groupclasses ON class_users.groupclass_id = groupclasses.id").
		Where("users.id = ?", userID).
		Scan(&userClassInfos).Error
	if err != nil {
		return nil, err
	}
	return userClassInfos, nil
}
