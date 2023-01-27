package database

import (
	"fmt"
	"microservice/models"
	"microservice/pkg/sql"
)

func RunMigration() {
	err := sql.DB.AutoMigrate(
		&models.Admin{},
		&models.Groupclass{},
		&models.SubClass{},
		&models.User{},
		&models.Journal{},
		&models.BillingMonthly{},
		&models.Notification{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
}
