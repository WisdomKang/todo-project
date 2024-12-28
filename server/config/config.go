package config

import (
	"log"

	"github.com/joho/godotenv"
)

const (
	EnviromentsPath string = "application.env"
)

func LoadEnv() {

	err := godotenv.Load("application.env")
	if err != nil {
		log.Fatalf("Could Not Load %v file", err)
	}
}
