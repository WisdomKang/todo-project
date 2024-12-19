package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("test hello!")

	dns := "host=localhost user=root password=todo5432 dbname=gorm port=15432 sslmode=disable TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
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
