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
	ticker := time.NewTicker(newsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			news, err := data.MarketNews()
			if err != nil{
				panic(err)
			}
			for _, n := range news {
				if err := t.Write(fmt.Sprintf("%s", n)); err != nil {
					panic(err)
				}
			}

		case <-ctx.Done():
			return
		}
	}
}
