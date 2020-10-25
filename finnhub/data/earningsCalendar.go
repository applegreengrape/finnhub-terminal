package data

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"
	"time"
)

// GetEarningsCals ..
func GetEarningsCals() (cals []string, pause bool, e error) {
	cfg := finnhub.NewSettingFromConfig()
	c := &client.Client{
		APIKey: cfg.APIKey,
	}
	path := "/calendar/earnings"

	now := time.Now()
	year, month, _ := now.Date()
	begin := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	from := fmt.Sprintf("%s", begin.Format("2020-01-02"))
	end := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location())
	to := fmt.Sprintf("%s", end.Format("2020-01-02"))

	for _, s := range cfg.Stocks {
		p := url.Values{}
		p.Add("from", from)
		p.Add("to", to)
		p.Add("symbol", s)
		res, err := c.FinnhubClient(finnhub.Version+path, p)
		if err != nil {
			return nil, true, err
		}
		var earningcals EarningsCalendarStruct
		json.NewDecoder(res.Body).Decode(&earningcals)

		if len(earningcals.EarningsCalendar) == 0 {
			str := fmt.Sprintf("🈳 no earning calendar info for %s from %s to %s. \n\n", s, from, to)
			cals = append(cals, str)
			pause = true
		} else {
			pause = false
			cal := earningcals.EarningsCalendar[0]
			meta := fmt.Sprintf("⏰ %s Q%d [%s]\n", cal.Date, cal.Quarter, cal.Symbol)
			cals = append(cals, meta)
			details := fmt.Sprintf("EpsActual: %.2f, EpsEstimate: %.2f, RevenueActual: %.2f, RevenueEstimate: %.2f \n\n", cal.EpsActual, cal.EpsEstimate, cal.RevenueActual, cal.RevenueEstimate)
			cals = append(cals, details)
		}
	}

	return cals, pause, nil
}
