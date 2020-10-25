package widgets

import (
	"context"
	"fmt"
	"time"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/mum4k/termdash/widgets/text"
)

// UpdateFinancialReports ..
func UpdateFinancialReports(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			d, err := data.GetFinancialReports()
			if err != nil {
				fmt.Println(err)
			}

			str := &strings.Builder{}
			table := tablewriter.NewWriter(str)
			table.SetHeader([]string{"filedDate", "type", "concept", "label", "value", "unit"})
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			table.SetAutoWrapText(true)
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
