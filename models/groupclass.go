package models

import (
	"time"
)

type Groupclass struct {
	Id         int        `json:"id" gorm:"PRIMARY_KEY"`
	Groupclass int        `json:"groupclass"`
	Startclass time.Time  `json:"startclass"`
	Endclass   time.Time  `json:"endclass"`
	Level      string     `json:"level"`
	Users      []User     `gorm:"many2many:class_users"`
	Subclass   []SubClass `gorm:"foreignkey:Groupclass"`
}

type SubClass struct {
	ID         int    `json:"id" gorm:"PRIMARY_KEY"`
	Groupclass int    `json:"groupclass"`
	Subclass   string `json:"subclass" gorm:"type: varchar(255)"`
	User       []User `gorm:"many2many:subclass_users"`
}

type ClassUser struct {
	GroupclassID uint
	UserID       uint
}

type SubclassUser struct {
	SubClassID uint
	UserID     uint
}
