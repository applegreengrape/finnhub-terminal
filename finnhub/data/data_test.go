package data

import (
	"fmt"
	"os"
	"testing"
	"github.com/olekukonko/tablewriter"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
)

func TestCompanyNews(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	n, err := CompanyNews()
	if err == nil {
		fmt.Println(n)
	}
}

func TestMarketNews(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	n, err := MarketNews()
	if err == nil {
		fmt.Println(n)
	}
}

func TestGetBasicFinancials(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	data, err := GetBasicFinancials()
	if err != nil {
		fmt.Println(err)
	}

	header := []string{"metric"}
	cfg := finnhub.NewSettingFromConfig()

	for _, s := range cfg.Stocks{
		header = append(header, s)
	}

	fmt.Println(header)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)                          
	table.Render()
}
