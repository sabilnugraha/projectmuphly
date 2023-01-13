package admindto

type AdminResponse struct {
	UserName string `gorm:"type: varchar(255)" json:"UserName"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
