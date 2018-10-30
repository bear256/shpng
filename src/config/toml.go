package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var conf config

func init() {
	if _, err := toml.DecodeFile("src/config/system.toml", &conf); err != nil {
		log.Fatalln(err)
	}
}

// Config method return the configuration with default toml file in dev stage
func Toml() config {
	return conf
}

// Load method return the configuration with the specified toml file
func LoadToml(fname string) config {
	if _, err := toml.DecodeFile(fname, &conf); err != nil {
		log.Fatalln(err)
	}
	return conf
}

type config struct {
	Storage storage
}

type storage struct {
	Name string
	Key  string
}
