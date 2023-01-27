package models

import (
	"time"
)

type Groupclass struct {
	Id         int        `json:"id" gorm:"PRIMARY_KEY"`
	Groupclass int        `json:"groupclass"`
	Startclass time.Time  `json:"startclass"`
	Endclass   time.Time  `json:"endclass"`
	User       []User     `gorm:"foreignkey:Classgroup"`
	Subclass   []SubClass `gorm:"foreignkey:Groupclass"`
	Status     string
}

type SubClass struct {
	ID         int    `json:"id" gorm:"PRIMARY_KEY"`
	Groupclass int    `json:"groupclass"`
	Subclass   string `json:"subclass" gorm:"type: varchar(255)"`
	User       []User `gorm:"foreignkey:Subclass"`
}
