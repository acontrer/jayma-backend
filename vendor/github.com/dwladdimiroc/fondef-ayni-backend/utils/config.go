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

type Files struct {
	Uri string `yaml:"uri"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Asertividad struct {
	Jar  string `yaml:"jar"`
	Main string `yaml:"main"`
}

type Configuration struct {
	Database    Database    `yaml:"database"`
	Files       Files       `yaml:"files"`
	Server      Server      `yaml:"server"`
	Asertividad Asertividad `yaml:"asertividad"`
}

func LoadConfig(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	Check(err)

	err = yaml.Unmarshal(bytes, &Config)
	Check(err)
}
