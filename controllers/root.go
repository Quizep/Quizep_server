package controllers

import "github.com/gin-gonic/gin"

func NewControllers(port string) {
	r := gin.Default()

	auth := r.Group("/auth")
	AuthContoller(auth)
}
