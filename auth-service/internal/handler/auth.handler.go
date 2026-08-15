package handler

import (
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	c.JSON(401, gin.H{"message": "auth handler"})
}

func LoginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "login handler"})
}
