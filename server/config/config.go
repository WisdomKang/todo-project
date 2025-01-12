package config

import (
	"fmt"
	"log"
	"os"
	"todo-project/models"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	EnviromentsPath string = "application.env"
)

func InitConfig() {
	LoadEnv()

	InitAuth()
	InitDatabase()
}

func LoadEnv() {
	err := godotenv.Load("./application.env")
	if err != nil {
		log.Fatalf("Could Not Load %v file", err)
	}
}

// Authentication Configuration
var GoogleAuthConfig *oauth2.Config

func InitAuth() {
	GoogleAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECERT"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		Scopes:       []string{"openid"},
		Endpoint:     google.Endpoint,
	}

	log.Printf("Oauth config %v", GoogleAuthConfig)

}

// Database Configuration
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

	Migration()
}

func Migration() {
	err := DB.AutoMigrate(&models.User{}, &models.ToDo{})

	if err != nil {
		log.Fatalf("Fail Migration : %v", err)
	}
}
