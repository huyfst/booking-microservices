package main

import (
	"booking-microservice/auth-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/auth/login", handler.LoginHandler)

	router.GET("/auth/verify", handler.AuthHandler)

	router.Run(":3000")
}
