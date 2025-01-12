package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-project/config"

	"github.com/gin-gonic/gin"
)

func GetOauthLoginPage(ctx *gin.Context) {
	url := config.GoogleAuthConfig.AuthCodeURL("")
	log.Printf("page redirect : %v", url)
	ctx.Redirect(http.StatusPermanentRedirect, url)
}

func ExchangeToken(ctx *gin.Context) {
	code := ctx.Query("code")

	token, err := config.GoogleAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("token exchange err : %v", err)
		return
	}

	client := config.GoogleAuthConfig.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("erro : %v", err)
		ctx.Error(err)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Printf("erro : %v", err)
		ctx.Error(err)
		return
	}

	log.Println("User Info:", userInfo)
}
