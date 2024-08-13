package services

import (
	"Quizep_server/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func User(c *gin.Context, db *gorm.DB) {
	var userDAO *entities.UserDAO
	var userDTO *entities.UserDTO

	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.Where("email = ?", userDTO.Email).First(&userDAO)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			newUser := &entities.UserDAO{
				UUID:
			}
		}
	}
}
