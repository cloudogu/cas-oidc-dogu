package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Issuer       string `yaml:"issuer"`
	FQDN         string `yaml:"fqdn"`
	Port         string `yaml:"port"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// ReadConfig reads the application configuration from a configuration file.
func ReadConfig(path string) (Config, error) {
	config := Config{}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return config, fmt.Errorf("could not find configuration at %s", path)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("failed to read configuration %s: %w", path, err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal configuration %s: %w", path, err)
	}

	return config, nil
}
