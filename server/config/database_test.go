package config_test

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(mysql.Open("todo_admin:todo5432@tcp(localhost:3306)/todo_database?charset=utf8"))
	if err != nil {
		log.Fatalf("DB connection error : %v", err)
	}

	code := m.Run()

	os.Exit(code)
}
