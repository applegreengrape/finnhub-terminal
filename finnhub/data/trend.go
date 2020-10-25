package data

import (
	"encoding/json"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"
)

// GetTrend ..
func GetTrend() (target []int, peers []int, e error) {
	cfg := finnhub.NewSettingFromConfig()
	c := &client.Client{
		APIKey: cfg.APIKey,
	}
	path := "/stock/recommendation"

	var b int
	var h int
	var sell int
	var sb int
	var ss int
	for _, s := range cfg.Stocks {
		param := url.Values{}
		param.Add("symbol", s)
		res, err := c.FinnhubClient(finnhub.Version+path, param)
		if err != nil {
			return nil, nil, err
		}
		var trend Trend
		json.NewDecoder(res.Body).Decode(&trend)

		if s == cfg.Stocks[0] {
			target = []int{trend[0].Buy, trend[0].Hold, trend[0].Sell, trend[0].StrongBuy, trend[0].StrongSell}
		}else{
			b += trend[0].Buy
			h += trend[0].Hold
			sell += trend[0].Sell
			sb += trend[0].StrongBuy
			ss += trend[0].StrongSell
		}
	}
	
	n := len(cfg.Stocks)-1
	peers = []int{b/n, h/n, sell/n, sb/n, ss/n}
	
	return target, peers, nil
}
