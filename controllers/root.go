package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewControllers(port string, db *gorm.DB) {
	r := gin.Default()

	auth := r.Group("/auth")
	AuthContoller(auth, db)

	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
