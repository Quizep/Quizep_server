package controllers

import (
	"Quizep_server/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthContoller(r *gin.RouterGroup, db *gorm.DB) {
	r.POST("/signup", func(c *gin.Context) {
		services.SignUp(c, db)
	})
}
