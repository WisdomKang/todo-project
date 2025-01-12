package route

import (
	"todo-project/auth"

	"github.com/gin-gonic/gin"
)

var Routes *gin.Engine

func SetRoute() {

	Routes = gin.Default()
	Protected := Routes.Group("/api")

	Protected.GET("/todo")
	Protected.POST("/todo")

	Routes.GET("/auth/v1/google", auth.GetOauthLoginPage)
	Routes.GET("/auth/v1/token", auth.ExchangeToken)

}
