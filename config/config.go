package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type AgsConfig struct {
	Connection     string `yaml:"connection"`
	BusinessId     string `yaml:"businessId"`
	CatalogId      string `yaml:"catalogId"`
	ApplicationId  string `yaml:"applicationId"`
	PrivateKeyPath string `yaml:"privateKeyPath"`
}

func (c *AgsConfig) GetConf() *AgsConfig {

	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
