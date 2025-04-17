package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

var configFile = os.Getenv("USER") + ".local/share/dfiles/config.json"

func ReadConfig() Config {
	byteValue, err := os.ReadFile(configFile)

	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func WriteConfig(config Config) {
	byteValue, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(configFile, byteValue, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
