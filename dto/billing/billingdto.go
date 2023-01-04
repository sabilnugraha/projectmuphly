package billingdto

import "time"

type BillingRequest struct {
	IdUser   int       `json:"user_id"`
	Month    time.Time `json:"month"`
	Category string    `json:"category"`
}
