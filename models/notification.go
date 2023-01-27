package models

type Notification struct {
	Id      int    `json:"Id" gorm:"PRIMARY_KEY"`
	UserID  int    `json:"user_id"`
	User    User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AdminID int    `json:"admin_id"`
	Admin   Admin  `json:"admin" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type    string `json:"type" gorm:"type: varchar(255)"`
	Status  string `json:"status" gorm:"type: varchar(255)"`
}
