package cmd

import (
	"fmt"
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	MySQL struct {
		Host string
	}
}

func NewCmd(filePath string) {
	c := new(Config)
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	fmt.Println(c.MySQL.Host)
	fmt.Println(c.Server.Port)
}
