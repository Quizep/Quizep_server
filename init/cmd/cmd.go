package cmd

import "os"

func NewCmd(filePath string) {
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}
}
