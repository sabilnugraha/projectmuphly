package journaldto

import "time"

type JournalRequest struct {
	Id            int       `json:"id" gorm:"PRIMARY_KEY"`
	IdTransaction int       `json:"transaction_id"`
	date          time.Time `json:"-"`
	Desc          string    `json:"desc"`
	OpsInput      *int      `json:"opsinput"`
	OpsOutput     int       `json:"opsoutput"`
	MonthlyInput  int       `json:"monthlyinput"`
	MonthlyOutput int       `json:"monthlyoutput"`
	MahadInput    int       `json:"mahadinput"`
	MahadOutput   int       `json:"mahadoutput"`
}
