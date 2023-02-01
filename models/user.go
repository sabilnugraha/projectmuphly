package models

import "time"

type User struct {
	Id         int          `json:"id" gorm:"PRIMARY_KEY"`
	Nis        int          `json:"nis"`
	Nisn       int          `json:"nisn"`
	Nik        int          `json:"nik"`
	FullName   string       `json:"fullName" gorm:"type: varchar(255)"`
	NickName   string       `json:"nickname" gorm:"type: varchar(255)"`
	Address    string       `json:"address" gorm:"type: varchar(255)"`
	BirthPlace string       `json:"birthplace" gorm:"type: varchar(255)"`
	BirthDate  time.Time    `json:"birthdate"`
	Phone      string       `json:"phone" gorm:"type: varchar(255)"`
	Photo      string       `json:"photo" gorm:"type: varchar(255)"`
	Category   string       `json:"category" gorm:"type: varchar(255)"`
	Status     string       `json:"status" gorm:"type: varchar(255)"`
	GroupClass []Groupclass `gorm:"many2many:class_users"`
	Subclass   []SubClass   `gorm:"many2many:subclass_users"`
	Password   string       `json:"Password" gorm:"type: varchar(255)"`
}

type UserClassDetails struct {
	FullName   string    `json:"fullName" gorm:"column:full_name"`
	StartClass time.Time `json:"startClass" gorm:"column:startclass"`
	EndClass   time.Time `json:"endClass" gorm:"column:endclass"`
}
