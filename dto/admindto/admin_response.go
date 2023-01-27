package admindto

type AdminResponse struct {
	UserName string `gorm:"type: varchar(255)" json:"UserName"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Position string `json:"Position" gorm:"type: varchar(255)"`
}

type CheckAuthResponse struct {
	ID       int    `json:"id"`
	UserName string `gorm:"type: varchar(255)" json:"UserName"`
}
