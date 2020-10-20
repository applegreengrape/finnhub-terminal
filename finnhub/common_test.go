package finnhub

import (
	"os"
	"fmt"
	"testing"
)

func TestFinnhubClient(t *testing.T) {
	os.Setenv("finnhub_config_path", "../default_config.json")
	s:= NewSettingFromConfig()

	testClient := &Client{
		symbols: s.symbols,
		apiKey: s.apiKey,
	}

	r, err := testClient.GetQuote()
	if err != nil {
		fmt.Println(err)
	}
	
	for _, stk := range r {
		fmt.Println(stk.Stock, stk.C)
	}

}

func TestWidget(t *testing.T) {
	l:= ListQuotes()
	
	fmt.Println(l.Rows)
}