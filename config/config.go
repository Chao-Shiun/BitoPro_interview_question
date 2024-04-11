package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	AppName string
	Port    string
	Logger  struct {
		Use         string
		Environment string
		LogLevel    string
		FileName    string
	}
}

func (c *Config) LoadConfig() *Config {
	path, err := filepath.Abs("config/config.yml")
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.LoadConfig()
	return cfg
}
