package main

import (
	"todo-project/config"
)

func main() {

	config.LoadEnv()
	config.InitDatabase()
	config.Migration()

}
