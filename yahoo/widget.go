package yahoo

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func dashboard(l *widgets.List, s *Config)(list *widgets.List){
	c := &Client{
        symbols: s.Stocks,
    }

    r, err := c.GetPrice()
    if err != nil {
        //
	}
	
    for idx, stk := range r {
        line := fmt.Sprintf("[%d], %s, %.2f", idx, stk.stock, stk.price)
        l.Rows = append(l.Rows, line)
    }
    l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	return l
}

// GetStockPrice ..
func GetStockPrice(){
	if err := ui.Init(); err != nil {
		fmt.Println(err)
	}
	defer ui.Close()

	s := NewSettingFromConfig()
	l := widgets.NewList()
	l.Title = "Stock Price"

	for _, view := range s.EnableViews {
		if view.Widget == "stock_price" {
			interval := time.Duration(view.Interval) * time.Second
			timer := time.NewTicker(interval)
			l.SetRect(view.WidgetRect.Top, view.WidgetRect.Left, view.WidgetRect.Width, view.WidgetRect.Heigt)
			ui.Render(l)
			previousKey := ""
			uiEvents := ui.PollEvents()
			for {
				select {
				case <- timer.C :
					e := <-uiEvents
					switch e.ID {
					case "q", "<C-c>":
						return
					case "j", "<Down>":
						l.ScrollDown()
					case "k", "<Up>":
						l.ScrollUp()
					case "<C-d>":
						l.ScrollHalfPageDown()
					case "<C-u>":
						l.ScrollHalfPageUp()
					case "<C-f>":
						l.ScrollPageDown()
					case "<C-b>":
						l.ScrollPageUp()
					case "g":
						if previousKey == "g" {
							l.ScrollTop()
						}
					case "<Home>":
						l.ScrollTop()
					case "G", "<End>":
						l.ScrollBottom()
					case "e":
						break
					}
					if previousKey == "g" {
						previousKey = ""
					} else {
						previousKey = e.ID
					}
					NewList := dashboard(l, s) 
					NewList.TextStyle = ui.NewStyle(ui.ColorYellow)
					NewList.WrapText = false
					ui.Render(NewList)
				}
			}
		}
	}
	
}
