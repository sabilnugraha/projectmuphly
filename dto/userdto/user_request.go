package userdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"FullName" validate:"required"`
	UserName string `gorm:"type: varchar(255)" json:"UserName" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"Password"`
}
