package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Token    string `yaml:"token"`
	Port     string `yaml:"port"`
	Services struct {
		Loms           string `yaml:"loms"`
		ProductService string `yaml:"productService"`
	} `yaml:"services"`
	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"dbname"`
	}
	MaxWorkers int `yaml:"max_workers"`
}

var ConfigData ConfigStruct

func Init() error {
	rawYAML, err := os.ReadFile("config.yml")
	if err != nil {
		return errors.WithMessage(err, "reading config file")
	}

	err = yaml.Unmarshal(rawYAML, &ConfigData)
	if err != nil {
		return errors.WithMessage(err, "parsing yaml")
	}

	return nil
}
