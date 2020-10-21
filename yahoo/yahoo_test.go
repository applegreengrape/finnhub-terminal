package yahoo

import (
	"fmt"
	"os"
	"testing"
)

func TestGetCurrentPrice(t *testing.T) {
	os.Setenv("finnhub_config_path", "../config.sample.json")
	stkPrice := GetCurrentPrice()

	fmt.Println(stkPrice)
}

func TestGetCurrentPriceChart(t *testing.T) {
	os.Setenv("finnhub_config_path", "../config.sample.json")
	prices := GetHistPrice()
	fmt.Println(prices[0])
}

func TestGetCurrentVolumeChart(t *testing.T) {
	os.Setenv("finnhub_config_path", "../config.sample.json")
	vol := GetHistVolume()
	fmt.Println(vol[0])

}
