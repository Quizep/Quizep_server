package reposistories

import (
	"Quizep_server/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dsn string) (*gorm.DB, error) {
	// 데이터베이스 연결 생성
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 데이터베이스 자동 마이그레이션
	err = db.AutoMigrate(&entities.UserDAO{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
