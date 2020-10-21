package widgets

import (
	"context"
	"fmt"
	"time"

	"github.com/applegreengrape/fintech-terminal/finnhub"
	"github.com/mum4k/termdash/widgets/text"
)

// RollingCompanyNews ..
func RollingCompanyNews(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			news := finnhub.CompanyNews()
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
