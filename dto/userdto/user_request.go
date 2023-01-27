package userdto

type RegisterRequest struct {
	FullName string `gorm:"type: varchar(255)" json:"FullName" validate:"required"`
	UserName string `gorm:"type: varchar(255)" json:"UserName" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"Password"`
}
type StudentAccountRequest struct {
	Nis      int    `json:"nis"`
	Password string `gorm:"type: varchar(255)" json:"Password"`
}

type StudentAccountResponse struct {
	Nis      int    `json:"nis"`
	Password string `gorm:"type: varchar(255)" json:"Password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
