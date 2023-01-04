package admindto

type AdminRequest struct {
	UserName string `gorm:"type: varchar(255)" json:"UserName" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"Password"`
}
