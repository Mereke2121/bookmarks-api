package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Port string `yaml:"server_port"`
	DB   struct {
		User    string `yaml:"user"`
		DBName  string `yaml:"dbname"`
		SSLMode string `yaml:"sslmode"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
	} `yaml:"postgres"`
}

func InitConfig() (*Config, error) {
	data, err := ioutil.ReadFile("./config/conf.yml")
	if err != nil {
		return nil, err
	}

	var conf *Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
