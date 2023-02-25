package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickchangch/at-assignment/service"
)

func AuthHandler(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]

	result, err := service.AuthService.ValidateJWT(tokenString)
	if !result || err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Next()
}
