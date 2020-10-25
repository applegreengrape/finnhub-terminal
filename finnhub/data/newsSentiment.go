package data

import (
	"encoding/json"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"
)

// GetNewsSentiment ..
func GetNewsSentiment() (t float64, p float64, e error) {
	cfg := finnhub.NewSettingFromConfig()
	c := &client.Client{
		APIKey: cfg.APIKey,
	}
	path := "/news-sentiment"

	var peers float64
	for _, s := range cfg.Stocks {
		param := url.Values{}
		param.Add("symbol", s)
		res, err := c.FinnhubClient(finnhub.Version+path, param)
		if err != nil {
			return 0, 0, err
		}
		var newsSentiment NewsSentiment
		json.NewDecoder(res.Body).Decode(&newsSentiment)

		if s == cfg.Stocks[0] {
			t = newsSentiment.Sentiment.BullishPercent
		}else{
			peers += newsSentiment.Sentiment.BullishPercent
		}
	}

	p = peers/float64(len(cfg.Stocks)-1)

	return t, p, nil
}
