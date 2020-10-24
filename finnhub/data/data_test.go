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

	n, p, err := CompanyNews()
	if err == nil {
		fmt.Println(p)
		for _, l := range n {
			fmt.Println(l)
		}
	}
}

func TestMarketNews(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	n, p, err := MarketNews()
	if err == nil {
		fmt.Println(n, p)
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


func TestGetEarningsCals(t *testing.T) {
	os.Setenv("finnhub_terminal_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	cals, p, err := GetEarningsCals()
	if err == nil {
		fmt.Println(p)

		for _, c := range cals {
			fmt.Println(c)
		}
	}
}