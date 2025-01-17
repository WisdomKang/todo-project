package auth

import (
	"log"
	"net/http"
	"todo-project/config"
	"todo-project/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func LoginPageHandler(ctx *gin.Context) {
	url := config.GoogleAuthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	log.Printf("page redirect : %v", url)
	ctx.Redirect(http.StatusPermanentRedirect, url)
}

func ExchangeTokenHandler(ctx *gin.Context) {
	code := ctx.Query("code")

	token, err := config.GoogleAuthConfig.Exchange(ctx, code)
	if err != nil {
		log.Printf("token exchange err : %v", err)
		return
	}

	utils.PrintStruct(token)

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error during token exchange"})
		return
	}

	ctx.SetCookie("id_token", idToken, 3600, "/", "", true, true)

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func RefreshTokenHandler(ctx *gin.Context) {

}
