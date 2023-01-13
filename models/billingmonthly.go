package models

import "time"

type BillingMonthly struct {
	Id          int       `json:"id" gorm:"PRIMARY_KEY"`
	UserID      int       `json:"user_id"`
	User        User      `json:"user"`
	Month       time.Time `json:"month"`
	PaymentDate time.Time `json:"paymentdate"`
	Ops         int       `json:"ops"`
	Monthly     int       `json:"monthly"`
	Dormitory   int       `json:"dormitory"`
	Status      string    `json:"status"`
}
