package models

import (
	"time"
)

type Journal struct {
	Id              int       `json:"id" gorm:"PRIMARY_KEY"`
	IdTransaction   int       `json:"transaction_id"`
	Date            time.Time `json:"date" gorm:"default:current_timestamp"`
	Desc            string    `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	OpsInput        int       `json:"opsinput" gorm:"default:0"`
	OpsOutput       int       `json:"opsoutput" gorm:"default:0"`
	MonthlyInput    int       `json:"monthlyinput" gorm:"default:0"`
	MonthlyOutput   int       `json:"monthlyoutput" gorm:"default:0"`
	MahadInput      int       `json:"mahadinput" gorm:"default:0"`
	MahadOutput     int       `json:"mahadoutput" gorm:"default:0"`
	DormitoryInput  int       `json:"dormitoryinput" gorm:"default:0"`
	DormitoryOutput int       `json:"dormitoryoutput" gorm:"default:0"`
	Uraian          string    `json:"uraian" gorm:"type: varchar(255)"`
}

type JournalView struct {
	Id            int       `json:"id" gorm:"PRIMARY_KEY"`
	IdTransaction int       `json:"transaction_id"`
	Date          time.Time `json:"date" gorm:"default:current_timestamp"`
	Desc          string    `json:"desc" form:"desc" gorm:"type: varchar(255)"`
	OpsInput      int       `json:"opsinput" gorm:"default:0"`
	OpsOutput     int       `json:"opsoutput" gorm:"default:0"`
	MonthlyInput  int       `json:"monthlyinput" gorm:"default:0"`
	MonthlyOutput int       `json:"monthlyoutput" gorm:"default:0"`
	MahadInput    int       `json:"mahadinput" gorm:"default:0"`
	MahadOutput   int       `json:"mahadoutput" gorm:"default:0"`
	Debet         int       `json:"debet"`
	Kredit        int       `json:"kredit"`
	Saldoakhir    int       `json:"saldoakhir"`
	Uraian        string    `json:"uraian" gorm:"type: varchar(255)"`
}
