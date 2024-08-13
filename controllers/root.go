package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewControllers(port string, sql *gorm.DB) error {
	r := gin.Default()

	auth := r.Group("/auth")
	AuthContoller(auth)

	err := r.Run(port)
	return err
}
