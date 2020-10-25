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

// CompanyNews ..
func CompanyNews() (companyNewsList []string, pause bool, e error) {
	cfg := finnhub.NewSettingFromConfig()
	c := &client.Client{
		APIKey: cfg.APIKey,
	}
	path := "/company-news"

	for _, s := range cfg.Stocks {

		p := url.Values{}
		p.Add("symbol", s)
		res, err := c.FinnhubClient(finnhub.Version+path, p)
		if err != nil {
			return nil, true, err
		}
		var companynews CompanyNewsStruct
		json.NewDecoder(res.Body).Decode(&companynews)

		if len(companynews) == 0 {
			str := fmt.Sprintf("🈳  no company news for %s yet.\n\n", s)
			companyNewsList = append(companyNewsList, str)
			pause = true

		} else {
			pause = false
			for _, n := range companynews {
				meta := fmt.Sprintf("📝 [%s] by %s (%s)\n", n.Related, n.Source, time.Unix(n.Datetime, 0).Format(time.RFC822Z))
				companyNewsList = append(companyNewsList, meta)

				if len(n.Headline) > 0 {
					hre := regexp.MustCompile("[[:^ascii:]]")
					h := hre.ReplaceAllLiteralString(n.Headline, "")
					headline := fmt.Sprintf("headline: %s\n", h)
					companyNewsList = append(companyNewsList, headline)
				}

				if len(n.Summary) > 0 {
					re := regexp.MustCompile("[[:^ascii:]]")
					s := re.ReplaceAllLiteralString(n.Summary, "")
					summary := fmt.Sprintf("summary: %s\n", s)
					companyNewsList = append(companyNewsList, summary)
				}
				companyNewsList = append(companyNewsList, "\n")
			}
		}
	}

	return companyNewsList, pause, nil
}
