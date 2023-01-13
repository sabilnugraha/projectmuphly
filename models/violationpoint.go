package models

type ViolationPoin struct {
	Id             int    `json:"id" gorm:"PRIMARY_KEY"`
	UserId         int    `json:"user_id"`
	Violation      string `json:"violation" gorm:"type: varchar(255)"`
	Point          int    `json:"poin"`
	Studentmanager string `json:"studentmanager" gorm:"type: varchar(255)"`
	Status         string `json:"status" gorm:"type: varchar(255)"`
}
