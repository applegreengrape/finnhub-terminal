package widgets

import (
	"context"
	"time"

	"github.com/applegreengrape/fintech-terminal/yahoo"
	"github.com/mum4k/termdash/widgets/sparkline"
)

// VolumeSparkChart ..
func VolumeSparkChart(ctx context.Context, sl *sparkline.SparkLine, delay time.Duration) {
	ticker := time.NewTicker(stockInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			vols := yahoo.GetHistVolume()
			if err := sl.Add(vols); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			return
		}
	}
}
