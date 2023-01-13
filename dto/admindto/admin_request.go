package admindto

type AdminRequest struct {
	UserName    string `gorm:"type: varchar(255)" json:"UserName" validate:"required"`
	Password    string `gorm:"type: varchar(255)" json:"Password"`
	Position    string `json:"Position" gorm:"type: varchar(255)"`
	BankAccount int    `json:"BankAccount"`
	BankName    string `json:"BankName" gorm:"type: varchar(255)"`
}
