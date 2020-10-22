package widgets

import (
	"context"
	"fmt"
	"time"

	"github.com/applegreengrape/finnhub-terminal/yahoo"
	"github.com/mum4k/termdash/widgets/text"
)

// UpdateStockPrice ..
func UpdateStockPrice(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(stockInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			str := fmt.Sprintf("🕐 : %s\n", now.Format(time.RFC3339))
			stocks := yahoo.GetCurrentPrice()
			for _, s := range stocks {
				str += fmt.Sprintf("%s", s)
			}

			if err := t.Write(str, text.WriteReplace()); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		}
	}
}
