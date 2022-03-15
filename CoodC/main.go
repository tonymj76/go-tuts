package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	var router = gin.Default()
	var credential Credential
	router.POST("/", func(ctx *gin.Context) {
		if err := ctx.ShouldBindJSON(&credential); err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		log.Println(credential)
		ctx.JSON(http.StatusOK, gin.H{"credential": credential})
	})
	router.Run()
}
