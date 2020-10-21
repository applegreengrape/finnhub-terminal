package finnhub

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config ..
type Config struct {
	IgAccountID   string `json:"ig_account_id"`
	AuthBearerTok string `json:"auth_bearer_tok"`
	IgAPIKey      string `json:"ig_api_key"`
	FinnhubAPIKey string `json:"finnhub_api_key"`
}

// NewSettingFromConfig - get the widgets setting from config.json
func NewSettingFromConfig() *Config {
	path, exists := os.LookupEnv("secret_json_path")
	if !exists {
		path = "secret.json"
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
