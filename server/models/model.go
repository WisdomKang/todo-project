package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `gorm:"unique"`
	Provider string
	ToDos    []ToDo `gorm:"foreignKey:UserID"`
}

type ToDo struct {
	gorm.Model

	UserID uint

	Contents   string
	DueDate    *time.Time
	Status     string
	Priority   string
	OrderIndex uint8
}
