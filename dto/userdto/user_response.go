package userdto

type RegisterResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	UserName string `gorm:"type: varchar(255)" json:"userName"`
}

type CheckAuthStudentResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"UserName"`
}
