package finnhub

import (
	"encoding/json"
	"fmt"
	"os"
)

// Settings defines the configuration properties for this module
type Settings struct {
	apiKey  string   `help:"Your finnhub API token."`
	symbols []string `help:"An array of stocks symbols (i.e. AAPL, MSFT)"`
}

// NewSettingFromConfig - get the widgets setting from config.json
func NewSettingFromConfig() *Settings {
	path, exists := os.LookupEnv("finnhub_config_path")
	if !exists {
		path = "default_config.json"
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

	setting := Settings{
		apiKey:  c.APIKey,
		symbols: c.Stocks,
	}

	return &setting
}
