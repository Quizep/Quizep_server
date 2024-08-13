package services

import (
	"Quizep_server/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func SignUp(c *gin.Context, db *gorm.DB) {
	var userDAO *entities.UserDAO
	var userDTO *entities.UserDTO

	if err := c.ShouldBind(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userDAO = &entities.UserDAO{
		UUID:     uuid.New().String(),
		Email:    userDTO.Email,
		Password: string(hasedPassword),
		Nickname: userDTO.Nickname,
	}

	if err := db.Create(&userDAO).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "성공",
		//"accessToken": ,
		//"refreshToken": ,
	})
}

func Login(c *gin.Context, db *gorm.DB) {

}
