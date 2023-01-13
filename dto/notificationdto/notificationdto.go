package notificationdto

type Notification struct {
	Id                     int    `json:"Id" gorm:"PRIMARY_KEY"`
	UserID                 int    `json:"user_id"`
	NotificationToPosition string `json:"notification_to_position" gorm:"type: varchar(255)"`
	AdminID                int    `json:"admin_id"`
	Type                   string `json:"type" gorm:"type: varchar(255)"`
	Status                 string `json:"status" gorm:"type: varchar(255)"`
}
