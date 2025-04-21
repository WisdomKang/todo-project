package main

import (
	"todo-project/config"
	route "todo-project/routes"
)

func main() {
	config.InitConfig()
	route.SetRoute()

	route.Routes.Run("localhost:9090")
}
