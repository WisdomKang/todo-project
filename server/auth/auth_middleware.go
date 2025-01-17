package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	idToken, err := ctx.Cookie("id_token")

	if err != nil {
		log.Printf("Authentication error %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		ctx.Abort()
		return
	}

	payload, err := VerifyIDToken(ctx, idToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
		ctx.Abort()
		return
	}

	email := payload["email"].(string)
	ctx.Set("userEmail", email)

	ctx.Next()
}
