package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	AddNotification(notif models.Notification) (models.Notification, error)
}

func RepositoryNotification(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddNotification(notif models.Notification) (models.Notification, error) {
	err := r.db.Create(&notif).Error

	return notif, err
}
