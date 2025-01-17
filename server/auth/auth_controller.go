package auth

import (
	"fmt"
	"net/http"
	"time"
	"todo-project/config"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")

	if err != nil || refreshToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing refresh tokne"})
		return
	}

	payload, err := idtoken.ParsePayload(refreshToken)

	exp := payload.Claims["ExpiresAt"].(int64)

	if exp > time.Now().Unix() {

	}

}

func VerifyIDToken(ctx *gin.Context, idToken string) (map[string]interface{}, error) {
	payload, err := idtoken.Validate(ctx, idToken, config.GoogleAuthConfig.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %v", err)
	}

	return payload.Claims, nil
}
