package yahoo

import (
	"encoding/json"
	"fmt"
)

// GetHistVolume ..
func GetHistVolume() (v []int) {
	cfg := NewSettingFromConfig()

	vol, err := cfg.getHistVolume(cfg.Stocks[0])
	if err != nil {
		fmt.Println(err)
	}
	return vol

}

/* -------------------- Unexported Functions -------------------- */

func (cfg *Config) getHistVolume(s string) ([]int, error) {
	v := []int{}

	resp, err := cfg.yahooHistRequest(s)
	if err != nil {
		return nil, err
	}

	var res YahooResp
	json.NewDecoder(resp.Body).Decode(&res)
	for _, vol := range res.Chart.Result[0].Indicators.Quote[0].Volume {
		v = append(v, vol)
	}

	return v, nil
}
