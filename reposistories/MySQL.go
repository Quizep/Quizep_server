package reposistories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return db, err
}
