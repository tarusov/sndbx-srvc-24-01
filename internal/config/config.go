package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	// HTTP server sonfig
	HTTPConfig struct {
		Port int `json:"port"`
	}

	// Global service config
	Config struct {
		HTTP HTTPConfig `json:"http"`
	}
)

// Read
func Read(fileName string) (*Config, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	c := &Config{}
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	return c, nil
}
