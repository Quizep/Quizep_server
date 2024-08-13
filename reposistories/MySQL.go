package reposistories

import (
	"Quizep_server/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	err = db.AutoMigrate(&entities.UserDAO{})
	return db, err
}
