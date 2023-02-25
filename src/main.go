package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rickchangch/at-assignment/base"
	"github.com/rickchangch/at-assignment/db"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	base.UrlMap(router)

	db.PostgreDB.Connect()

	router.Run(":80")
}
