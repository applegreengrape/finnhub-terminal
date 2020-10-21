package yahoo

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewSettingFromConfig - get the widgets setting from config.json
func NewSettingFromConfig() *Config {
	path, exists := os.LookupEnv("finnhub_config_path")
	if !exists {
		path = "config.json"
	}

	config, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var c Config
	err = json.NewDecoder(config).Decode(&c)
	if err != nil {
		fmt.Println(err)
	}

	return &c
}
