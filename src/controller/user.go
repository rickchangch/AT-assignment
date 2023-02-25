package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

var UserController User

func (u *User) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "Success"})
}
