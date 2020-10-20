package finnhub

import (
	"fmt"
	"github.com/gizak/termui/v3/widgets"
)

// ListQuotes ..
func ListQuotes() (list *widgets.List){
	l := widgets.NewList()
	l.Title = "Stock Price"
	s:= NewSettingFromConfig()
	
	c := &Client{
		symbols: s.symbols,
		apiKey: s.apiKey,
	}

	r, err := c.GetQuote()
	if err != nil {
		//
	}

	l.Rows = []string{}

	for idx, stk := range r {
		line := fmt.Sprintf("[%d], %s, %.2f", idx, stk.Stock, stk.C)
		l.Rows = append(l.Rows, line)
	}
	l.SetRect(0, 0, 50, 10)

	return l
}