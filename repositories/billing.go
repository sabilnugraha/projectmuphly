package repositories

import (
	"microservice/models"

	"gorm.io/gorm"
)

type BillingRepository interface {
	AddBilling(product models.BillingMonthly) (models.BillingMonthly, error)
	GetUserClassInfoByUserID(userID int) (models.UserClassDetails, error)
}

func RepositoryBilling(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddBilling(billing models.BillingMonthly) (models.BillingMonthly, error) {
	err := r.db.Create(&billing).Error

	return billing, err
}

func (r *repository) GetUserClassInfoByUserID(userID int) (models.UserClassDetails, error) {
	var userClassInfos models.UserClassDetails
	err := r.db.Table("class_users").
		Select("users.full_name, groupclasses.startclass, groupclasses.endclass").
		Joins("JOIN users ON class_users.user_id = users.id").
		Joins("JOIN groupclasses ON class_users.groupclass_id = groupclasses.id").
		Where("users.id = ?", userID).
		Scan(&userClassInfos).Error

	return userClassInfos, err
}
