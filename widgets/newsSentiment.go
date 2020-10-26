package widgets

import (
	"context"
	"time"

	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/mum4k/termdash/widgets/donut"
)

// UpdateSentimentDonut ..
func UpdateSentimentDonut(ctx context.Context, d *donut.Donut, which string) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			target, peer, err := data.GetNewsSentiment()
			if err != nil {
				panic(err)
			}

			if which == "target" {
				if err := d.Percent(int(target * 100) ); err != nil {
					panic(err)
				}
			}

			if which == "peers" {
				if err := d.Percent(int(peer * 100) ); err != nil {
					panic(err)
				}
			}

			time.Sleep(finnhubSentimentInterval)

		case <-ctx.Done():
			return
		}
	}
}
