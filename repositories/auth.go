package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AddAdmin(admin models.Admin) (models.Admin, error)
	CreateStudentAccount(student models.User) (models.User, error)
	Login(username string) (models.Admin, error)
	LoginStudent(Nis int) (models.User, error)
	GetAdmin(ID int) (models.Admin, error)
	GetStudentId(ID int) (models.User, error)
	FindLastNis() (int, error)
	UpdateStatusByUserId(userId int, status string) error
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddAdmin(user models.Admin) (models.Admin, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) CreateStudentAccount(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Login(username string) (models.Admin, error) {
	var user models.Admin
	err := r.db.First(&user, "user_name=?", username).Error

	return user, err
}

func (r *repository) LoginStudent(Nis int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "nik=?", Nis).Error

	return user, err
}

func (r *repository) GetAdmin(ID int) (models.Admin, error) {
	var user models.Admin
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) GetStudentId(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) FindLastNis() (int, error) {
	var user models.User

	err := r.db.Where("nis = (SELECT MAX(nis) FROM users)").First(&user).Error

	return user.Nis, err
}

func (r *repository) UpdateStatusByUserId(userId int, status string) error {
	return r.db.Model(&models.Notification{}).Where("user_id = ?", userId).Update("status", status).Error
}
