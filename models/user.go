package models

type User struct {
	Id       int    `json:"id" gorm:"PRIMARY_KEY"`
	Nis      int    `json:"nis"`
	FullName string `json:"fullName" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Category string `json:"category" gorm:"type: varchar(255)"`
	Status   string `json:"status" gorm:"type: varchar(255)"`
	Angkatan int    `json:"angkatan"`
}
