package utils

import (
	"gopkg.in/yaml.v2"

	"io/ioutil"
)

var (
	Config Configuration
)

type Database struct {
	Rdbms   string `yaml:"rdbms"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Ip      string `yaml:"ip"`
	Port    string `yaml:"port"`
	Name    string `yaml:"name"`
	Sslmode string `yaml:"sslmode"`
}


type Server struct {
	Port string `yaml:port`
}

type Configuration struct {
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

func LoadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(bytes, &Config)
	Check(err)
}
