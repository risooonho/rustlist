package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

var Config config

type config struct {
	App appConfig
}

type appConfig struct {
	Addr string
}

func LoadConfig(path string) {
	rawData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading config file\n", err)
	}

	_, err = toml.Decode(string(rawData), &Config)
	if err != nil {
		log.Fatal("Error reading config file\n", err)
	}
}
