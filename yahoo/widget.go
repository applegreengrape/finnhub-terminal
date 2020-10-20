package yahoo

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

// GetStockPrice ..
func GetStockPrice()(list *widgets.List){
	s := NewSettingFromConfig()
	c := &Client{
		symbols: s.Stocks,
	}

	r, err := c.GetPrice()
	if err != nil {}
	
	l := widgets.NewList()
	l.Title = "Stock Price"
	l.Rows = []string{}

	for idx, stk := range r {
        line := fmt.Sprintf("[%d], %s, %.2f", idx, stk.stock, stk.price)
        l.Rows = append(l.Rows, line)
	}
	
	return l
}
