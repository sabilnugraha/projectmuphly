package models

type PaymentList struct {
	ID         int            `json:"id" gorm:"PRIMARY_KEY"`
	Billing_Id int            `json:"billing_id"`
	Billing    BillingMonthly `json:"billing_monthly" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID     int            `json:"user_id"`
	User       User           `json:"user"`
	SubTotal   *int           `json:"subtotal"`
}
