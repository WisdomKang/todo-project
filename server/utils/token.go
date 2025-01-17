package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func CreateTokenSource(ctx *gin.Context, config *oauth2.Config, refreshToken string) oauth2.TokenSource {
	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	return config.TokenSource(ctx, token)

}

func GetNewAccessToken(ctx *gin.Context, tokenSource oauth2.TokenSource) (*oauth2.Token, error) {
	token, err := tokenSource.Token()

	if err != nil {
		log.Printf("Error refrehsin token: %v", err)
		return nil, err
	}

	return token, nil
}
