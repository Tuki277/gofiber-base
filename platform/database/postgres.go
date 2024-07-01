package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/app/models"
	"main.go/pkg/utils"
)

var DB *gorm.DB

func PostgreSQLConnection() {
	postgresConnURL, _ := utils.ConnectionURLBuilder("pgx")
	databaseConnection, err := gorm.Open(postgres.Open(postgresConnURL))

	if err != nil {
		panic("database connect failed")
	}
	DB = databaseConnection
	fmt.Println("database connected")

	AutoMigrate(databaseConnection)
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		models.Base{},
	)
}
