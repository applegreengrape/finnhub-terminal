package widgets

import (
	"context"
	"fmt"
	"time"

	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/mum4k/termdash/widgets/text"
)

// RollingEarningsCalendar ..
func RollingEarningsCalendar(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cals, pause, err := data.GetEarningsCals()
			if err != nil {
				panic(err)
			}
			for _, cal := range cals {
				if err := t.Write(fmt.Sprintf("%s", cal)); err != nil {
					panic(err)
				}
			}

			if pause {
				time.Sleep(finnhubPause)
			}

			time.Sleep(finnhubEarningCals)

		case <-ctx.Done():
			return
		}
	}
}
