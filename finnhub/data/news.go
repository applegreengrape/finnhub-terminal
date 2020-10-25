package data

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"
	"regexp"
	"time"
)

// MarketNews ..
func MarketNews() (mktNewsList []string, pause bool, e error) {
	cfg := finnhub.NewSettingFromConfig()
	c := &client.Client{
		APIKey: cfg.APIKey,
	}
	path := "/news"

	p := url.Values{}
	res, err := c.FinnhubClient(finnhub.Version+path, p)
	if err != nil {
		return nil, true, err
	}

	var mktnews News
	json.NewDecoder(res.Body).Decode(&mktnews)

	if len(mktnews) == 0 {
		str := "🈳 no market news for yet.\n"
		mktNewsList = append(mktNewsList, str)
		pause = true
	} else {
		pause = false
		for _, n := range mktnews {
			meta := fmt.Sprintf("🗞️  [%s] by %s (%s)\n", n.Category, n.Source, time.Unix(n.Datetime, 0).Format(time.RFC822Z))
			mktNewsList = append(mktNewsList, meta)

			if len(n.Headline) > 0 {
				hre := regexp.MustCompile("[[:^ascii:]]")
				h := hre.ReplaceAllLiteralString(n.Headline, "")
				headline := fmt.Sprintf("headline: %s\n", h)
				mktNewsList = append(mktNewsList, headline)
			}

			if len(n.Summary) > 0 {
				re := regexp.MustCompile("[[:^ascii:]]")
				s := re.ReplaceAllLiteralString(n.Summary, "")
				summary := fmt.Sprintf("summary: %s\n", s)
				mktNewsList = append(mktNewsList, summary)
			}

			mktNewsList = append(mktNewsList, "\n")
		}
	}

	return mktNewsList, pause, nil
}
