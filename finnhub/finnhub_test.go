package finnhub

import (
	"fmt"
	"os"
	"testing"
)

func TestMarketNews(t *testing.T) {
	os.Setenv("secret_json_path", "../secret.json")

	news := MarketNews()
	fmt.Println(news)
}

func TestCompanyNews(t *testing.T) {
	os.Setenv("secret_json_path", "../secret.json")

	news := CompanyNews()
	fmt.Println(news)
}
