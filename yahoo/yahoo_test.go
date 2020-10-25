package yahoo

import (
	"fmt"
	"os"
	"testing"
)

func TestGetCurrentPrice(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	stkPrice := GetCurrentPrice()

	fmt.Println(stkPrice)
}

func TestGetCurrentPriceChart(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	prices := GetHistPrice()
	fmt.Println(prices[0])
}

func TestGetCurrentVolumeChart(t *testing.T) {
	os.Setenv("finnhub_config", "/Users/pingzhouliu/Documents/playground/finnhub-terminal/config.json")
	vol := GetHistVolume()
	fmt.Println(vol[0])

}
