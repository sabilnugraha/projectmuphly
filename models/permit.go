package models

import "time"

type ShortTermPermit struct {
	Id            int       `json:"id" gorm:"PRIMARY_KEY"`
	UserId        int       `json:"user_id"`
	Permittime    time.Time `json:"permittime"`
	Backtime      time.Time `json:"backtime"`
	Reason        string    `json:"reason" gorm:"type: varchar(25time)"`
	Status        string    `json:"status" gorm:"type: varchar(255)"`
	Permitmanager string    `json:"permitmanager" gorm:"type: varchar(255)"`
}

type LongTermPermit struct {
	Id            int       `json:"id" gorm:"PRIMARY_KEY"`
	UserId        int       `json:"user_id"`
	PermitDate    time.Time `json:"permit_date"`
	BackDate      time.Time `json:"back_date"`
	Reason        string    `json:"reason" gorm:"type: varchar(255)"`
	Status        string    `json:"status" gorm:"type: varchar(255)"`
	Permitmanager string    `json:"permitmanager" gorm:"type: varchar(255)"`
}

type LongTermPermitRequest struct {
	Id            int       `json:"id" gorm:"PRIMARY_KEY"`
	UserId        int       `json:"user_id"`
	PermitDate    time.Time `json:"permit_date"`
	BackDate      time.Time `json:"back_date"`
	Reason        string    `json:"reason" gorm:"type: varchar(255)"`
	Status        string    `json:"status" gorm:"type: varchar(255)"`
	Permitmanager string    `json:"permitmanager" gorm:"type: varchar(255)"`
}
