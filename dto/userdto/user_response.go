package userdto

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	UserName string `gorm:"type: varchar(255)" json:"userName"`
}
