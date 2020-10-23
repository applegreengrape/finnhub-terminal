package data

import (
	"fmt"
	"os"
	"testing"
)

func TestCompanyNews(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/fintech-terminal/config.json")

	n, err := CompanyNews()
	if err == nil {
		fmt.Println(n)
	}
}

func TestMarketNews(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/fintech-terminal/config.json")

	n, err := MarketNews()
	if err == nil {
		fmt.Println(n)
	}
}
