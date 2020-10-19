package finnhub

import (
	"fmt"
	"testing"
)

func TestFinnhubClient(t *testing.T) {
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