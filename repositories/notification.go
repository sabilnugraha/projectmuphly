package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	AddNotification(notif models.Notification) (models.Notification, error)
	GetAllAdminByPosition(position string) ([]models.Admin, error)
	GetNotification(admin int, types string, status string) ([]models.Notification, error)
	FindNotificationId(UserId int) ([]models.Notification, error)
}

func RepositoryNotification(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddNotification(notif models.Notification) (models.Notification, error) {
	err := r.db.Create(&notif).Error

	return notif, err
}

func (r *repository) GetNotification(admin int, types string, status string) ([]models.Notification, error) {
	var notif []models.Notification
	err := r.db.Where("admin_id = ? AND type = ? AND status = ?", admin, types, status).Select("*").Find(&notif).Error

	return notif, err
}

func (r *repository) GetAllAdminByPosition(position string) ([]models.Admin, error) {
	var admin []models.Admin
	err := r.db.Where("position = ?", position).Select("*").Find(&admin).Error

	return admin, err
}

func (r *repository) FindNotificationId(UserID int) ([]models.Notification, error) {
	var notif []models.Notification
	err := r.db.Preload("User").Find(&notif, "admin_id = ?", UserID).Find(&notif, "status", "todo").Error

	return notif, err
}
