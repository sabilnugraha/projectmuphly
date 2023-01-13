package studentdto

type AddPhoto struct {
	Photo string `json:"photo" gorm:"type: varchar(255)"`
}
