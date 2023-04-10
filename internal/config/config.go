package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	Database Database `yaml:"database"`
	Http     Http     `yaml:"http"`
}

type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	Driver     string `yaml:"driver"`
}

type Http struct {
	Port string `yaml:"port"`
}

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}
	return Config{}
}

func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&cfg)
	if err != nil {
		return err
	}

	return nil
}
