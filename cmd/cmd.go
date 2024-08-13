package cmd

import (
	"Quizep_server/controllers"
	"Quizep_server/reposistories"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	MySQL struct {
		Host string
		Jwt  string
	}
}

func NewCmd(filePath string) {
	c := new(Config)
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	db := reposistories.NewRepositories(c.MySQL.Host)
	controllers.NewControllers(c.Server.Port, db)
}
