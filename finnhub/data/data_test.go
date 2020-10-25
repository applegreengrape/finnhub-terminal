package data

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/olekukonko/tablewriter"
)

func TestCompanyNews(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	n, p, err := CompanyNews()
	if err == nil {
		fmt.Println(p)
		for _, l := range n {
			fmt.Println(l)
		}
	}
}

func TestMarketNews(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	n, p, err := MarketNews()
	if err == nil {
		fmt.Println(n, p)
	}
}

func TestGetBasicFinancials(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	data, err := GetBasicFinancials()
	if err != nil {
		fmt.Println(err)
	}

	header := []string{"metric"}
	cfg := finnhub.NewSettingFromConfig()

	for _, s := range cfg.Stocks {
		header = append(header, s)
	}

	fmt.Println(header)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
	fmt.Println(tableString.String())

}

func TestGetEarningsCals(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")

	cals, p, err := GetEarningsCals()
	if err == nil {
		fmt.Println(p)

		for _, c := range cals {
			fmt.Println(c)
		}
	}
}

func TestGetFinancialReports(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	data, err := GetFinancialReports()
	if err != nil {
		fmt.Println(err)
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"filedDate", "type", "concept", "label", "value", "unit"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.AppendBulk(data)
	table.Render()
	fmt.Println(tableString.String())
}
