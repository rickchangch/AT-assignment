package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickchangch/at-assignment/db"
	"github.com/rickchangch/at-assignment/model"
	"github.com/rickchangch/at-assignment/service"
)

type Auth struct{}

var AuthController Auth

// Sign up a new user
func (a *Auth) SignUp(c *gin.Context) {

	jsonBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// @todo hash pwd
	user := model.User{}
	err = json.Unmarshal(jsonBody, &user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	tx := db.PostgreDB.Conn.Create(user)
	if tx.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Sign in
func (a *Auth) SignIn(c *gin.Context) {
	jsonBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := model.User{}
	err = json.Unmarshal(jsonBody, &user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var result model.User
	tx := db.PostgreDB.Conn.
		Where("acct = ? AND pwd = ?", user.Acct, user.Pwd).
		Find(&result)
	if tx.Error != nil || result.Acct == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, "Account or password not found.")
		return
	}

	token, err := service.AuthService.GenerateJWT(result.Acct)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
