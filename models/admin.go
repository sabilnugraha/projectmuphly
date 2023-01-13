package models

type Admin struct {
	Id            int            `json:"Id" gorm:"PRIMARY_KEY"`
	UserName      string         `json:"UserName" gorm:"type: varchar(255)"`
	Password      string         `json:"Password" gorm:"type: varchar(255)"`
	Position      string         `json:"Position" gorm:"type: varchar(255)"`
	BankAccount   int            `json:"BankAccount"`
	BankName      string         `json:"BankName" gorm:"type: varchar(255)"`
	Notifications []Notification `json:"notifications"`
}
