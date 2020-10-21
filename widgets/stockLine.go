package widgets

import (
	"context"
	"time"

	"github.com/applegreengrape/fintech-terminal/yahoo"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/linechart"
)

// StockLineChart ..
func StockLineChart(ctx context.Context, lc *linechart.LineChart, delay time.Duration) {
	ticker := time.NewTicker(stockInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			inputs := yahoo.GetHistPrice()
			cfg := yahoo.NewSettingFromConfig()
			if err := lc.Series(cfg.Stocks[0], inputs,
				linechart.SeriesCellOpts(cell.FgColor(cell.ColorYellow)),
			); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}
