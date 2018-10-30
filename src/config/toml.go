package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var conf Config

func init() {
	if _, err := toml.DecodeFile("src/config/system.toml", &conf); err != nil {
		log.Fatalln(err)
	}
}

// Toml method return the configuration with default toml file in dev stage
func Toml() Config {
	return conf
}

// LoadToml method return the configuration with the specified toml file
func LoadToml(fname string) Config {
	if _, err := toml.DecodeFile(fname, &conf); err != nil {
		log.Fatalln(err)
	}
	return conf
}

// Config is defined for global configuration
type Config struct {
	Storage storage
}

type storage struct {
	Name string
	Key  string
}
