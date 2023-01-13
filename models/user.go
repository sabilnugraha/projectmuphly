package models

import "time"

type User struct {
	Id            int            `json:"id" gorm:"PRIMARY_KEY"`
	Nis           int            `json:"nis"`
	Nisn          int            `json:"nisn"`
	Nik           int            `json:"nik"`
	FullName      string         `json:"fullName" gorm:"type: varchar(255)"`
	NickName      string         `json:"nickname" gorm:"type: varchar(255)"`
	Address       string         `json:"address" gorm:"type: varchar(255)"`
	BirthPlace    string         `json:"birthplace" gorm:"type: varchar(255)"`
	BirthDate     time.Time      `json:"birthdate" gorm:"type: varchar(255)"`
	Phone         string         `json:"phone" gorm:"type: varchar(255)"`
	Photo         string         `json:"photo" gorm:"type: varchar(255)"`
	Category      string         `json:"category" gorm:"type: varchar(255)"`
	Status        string         `json:"status" gorm:"type: varchar(255)"`
	Angkatan      int            `json:"angkatan"`
	Notifications []Notification `json:"notifications"`
}
