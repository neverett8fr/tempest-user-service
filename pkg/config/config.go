package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	route = "config/config.yaml"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Service Service `yaml:"user-service-config"`
}

func Initialise() *Config {

	conf := Config{}

	yamlFile, err := ioutil.ReadFile(route)
	if err != nil {
		log.Fatalf("issue finding config yaml, err %v", err)
		return &Config{}
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("issue unmarshalling config yaml, err %v", err)
		return &Config{}
	}

	return &conf
}
