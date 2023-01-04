package admindto

type AdminResponse struct {
	UserName string `gorm:"type: varchar(255)" json:"UserName"`
}
