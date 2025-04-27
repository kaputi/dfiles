package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

func newConfig() Config {
	return Config{
		Include: []string{},
		Exclude: []string{},
	}
}

var configFile = os.Getenv("HOME") + "/.local/share/dfiles/config.json"

func ReadConfig() Config {
	byteValue, err := os.ReadFile(configFile)

	if err != nil {
		return newConfig()
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	return config
}

func WriteConfig(config Config) {
	dir := filepath.Dir(configFile)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("Error creating directories: %v", err)
	}

	byteValue, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Error encoding config file: %v", err)
	}
	err = os.WriteFile(configFile, byteValue, 0644)
	if err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}
}
