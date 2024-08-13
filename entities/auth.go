package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
)

type StringSlice []string

func (ss *StringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal StringSlice value")
	}

	return json.Unmarshal(bytes, ss)
}

func (ss StringSlice) Value() (driver.Value, error) {
	return json.Marshal(ss)
}

type UserDAO struct {
	gorm.Model
	UUID       string `gorm:"primary_key"`
	Email      string `gorm:"unique"`
	Nickname   string
	Sex        string
	Password   string
	CreateExam StringSlice
	SolvedExam StringSlice
}

type UserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickaname"`
	Sex      string `json:"sex"`
}
