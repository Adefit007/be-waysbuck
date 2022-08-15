package database

import (
	"fmt"
	"waysbuck/models"
	"waysbuck/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Profile{},
		&models.Topping{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
