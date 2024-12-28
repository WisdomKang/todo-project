package config

import (
	"fmt"
	"log"
	"os"
	"todo-project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DbName     = "DB_NAME"
	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
	DbHost     = "DB_HOST"
	DbPort     = "DB_PORT"
)

var DB *gorm.DB

func InitDatabase() {
	dbName := os.Getenv(DbName)
	dbUser := os.Getenv(DbUser)
	dbPassword := os.Getenv(DbPassword)
	dbHost := os.Getenv(DbHost)
	dbPort := os.Getenv(DbPort)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("Database conntection failed : %v", err)
	}
}

func Migration() {
	err := DB.AutoMigrate(&models.User{}, &models.ToDo{})

	if err != nil {
		log.Fatalf("Fail Migration : %v", err)
	}
}
