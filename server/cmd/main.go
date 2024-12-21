package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("test hello!")

	dns := "todo_admin:todo5432@tcp(127.0.0.1:3306)/todo_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	db.AutoMigrate(&Test{})
}

type Test struct {
	gorm.Model

	Name string
	Agwe int
}
