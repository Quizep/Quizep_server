package services

import (
	"Quizep_server/init/cmd"
	"time"
)

var (
	jwtSecret       []byte
	accessTokenExp  = time.Hour
	refreshTokenExp = time.Hour * 24 * 7
)

func init() {
	config := cmd.GetConfig()
	jwtSecret = []byte(config.MySQL.Jwt)
}

func generateTokenPair() {
	
}
