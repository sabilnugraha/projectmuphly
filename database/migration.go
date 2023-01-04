package database

import (
	"fmt"
	"microservice/models"
	"microservice/pkg/sql"
)

func RunMigration() {
	err := sql.DB.AutoMigrate(
		&models.Admin{},
		&models.User{},
		&models.Journal{},
		&models.BillingMonthly{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
}
