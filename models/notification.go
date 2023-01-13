package models

type Notification struct {
	Id      int    `json:"Id" gorm:"PRIMARY_KEY"`
	UserID  int    `json:"user"`
	AdminID int    `json:"admin"`
	Type    string `json:"type" gorm:"type: varchar(255)"`
	Status  string `json:"status" gorm:"type: varchar(255)"`
}
