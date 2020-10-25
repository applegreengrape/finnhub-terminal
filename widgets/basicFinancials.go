package widgets

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/mum4k/termdash/widgets/text"
	"github.com/olekukonko/tablewriter"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
)

// UpdateBasicFinancials ..
func UpdateBasicFinancials(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			d, err := data.GetBasicFinancials()
			if err != nil {
				fmt.Println(err)
			}
		
			header := []string{"metric"}
			cfg := finnhub.NewSettingFromConfig()
		
			for _, s := range cfg.Stocks{
				header = append(header, s)
			}

			str := &strings.Builder{}
			table := tablewriter.NewWriter(str)
			table.SetHeader(header)
			table.AppendBulk(d)
			table.Render()

			if err := t.Write(str.String(), text.WriteReplace()); err != nil {
				panic(err)
			}

			time.Sleep(finnhubFundamentalInterval)
			
		case <-ctx.Done():
			return
		}
	}
}
