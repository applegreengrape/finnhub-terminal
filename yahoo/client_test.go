package yahoo

import (
	"fmt"
	"os"
	"testing"
)

func TestFinnhubClient(t *testing.T) {
	os.Setenv("finnhub_config_path", "../config.json")
	s := NewSettingFromConfig()

	testClient := &Client{
		symbols: s.Stocks,
	}

	r, err := testClient.GetPrice()
	if err != nil {
		fmt.Println(err)
	}

	for _, stk := range r {
		fmt.Println(stk.stock, stk.price)
	}

}
