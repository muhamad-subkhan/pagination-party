package database

import (
	"fmt"
	"party/models"
	"party/pkg/mysql"
)

func Migration() {
	err := mysql.DB.AutoMigrate(
		&models.Profile{},
		&models.Product{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
	fmt.Println("Migration Succes")
}