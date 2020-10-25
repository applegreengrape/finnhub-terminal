package widgets

import (
	"context"
	"fmt"
	"time"

	"github.com/applegreengrape/finnhub-terminal/finnhub/data"
	"github.com/mum4k/termdash/widgets/text"
)

// RollingCompanyNews ..
func RollingCompanyNews(ctx context.Context, t *text.Text) {
	ticker := time.NewTicker(finnhubInit)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			companyNews, pause, err := data.CompanyNews()
			if err != nil {
				panic(err)
			}
			for _, n := range companyNews {
				if err := t.Write(fmt.Sprintf("%s", n)); err != nil {
					panic(err)
				}
			}

			if pause {
				time.Sleep(finnhubPause)
			}

			time.Sleep(finnhubCN)

		case <-ctx.Done():
			return
		}
	}
}
