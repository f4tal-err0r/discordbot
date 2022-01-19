package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	token	string	`yaml:"token"`
	YTToken string  `yaml:"YTToken`
}

func NewConf(path string) (*Config, error) {
	//init config
	config := &Config{}

	if err := PathValidator(path); err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParseFlags() (string, error) {
	var path string
	flag.StringVar(&path, "config", "./config.yml", "path to config file")
	if err := PathValidator(path); err != nil {
		return "", err
	}

	return path, nil
}

func PathValidator(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("%s is a dir, not a file.", path)
	}
}