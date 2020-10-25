package widgets

import (
	"context"
	"fmt"
	"time"

	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/mum4k/termdash/widgets/text"
)

// RollingNews ..
func RollingNews(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			news, pause, err := data.MarketNews()
			if err != nil {
				panic(err)
			}
			for _, n := range news {
				if err := t.Write(fmt.Sprintf("%s", n)); err != nil {
					panic(err)
				}
			}
			
			if pause {
				time.Sleep(finnhubPause)
			}

			time.Sleep(finnhubInterval)

		case <-ctx.Done():
			return
		}
	}
}
