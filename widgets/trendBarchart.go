package widgets

import (
	"context"
	"time"

	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
)

// UpdateTrendBarChart ..
func UpdateTrendBarChart(ctx context.Context, bc *barchart.BarChart, which string) {
	const max = 5

	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			target, peers, err := data.GetTrend()
			if err != nil {
				panic(err)
			}
			if which == "target" {
				if err := bc.Values(target, max); err != nil {
					panic(err)
				}
			}
			if which == "peers" {
				if err := bc.Values(peers, max); err != nil {
					panic(err)
				}
			}

			time.Sleep(finnhubSentimentInterval)

		case <-ctx.Done():
			return
		}
	}
}