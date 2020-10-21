package finnhub

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/fintech-terminal/yahoo"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// CompanyNews ..
func CompanyNews() []string {
	cfg := NewSettingFromConfig()

	news, err := cfg.getCompanyNews()
	if err != nil {
		//
	}

	return news
}

/* -------------------- Unexported Functions -------------------- */

var (
	finnhubCompNewsURL = &url.URL{Scheme: "https", Host: "finnhub.io", Path: "/api/v1/company-news"}
)

func (cfg *Config) getCompanyNews() ([]string, error) {

	companyNewsList := []string{}
	stocks := yahoo.NewSettingFromConfig()

	for _, s := range stocks.Stocks {
		resp, err := cfg.finnhubResCompanyNews(s)
		if err != nil {
			return nil, err
		}
		var companynews CompanyNewsStruct
		json.NewDecoder(resp.Body).Decode(&companynews)

		for _, n := range companynews {
			meta := fmt.Sprintf("🗞️  [%s] by %s (%s)\n", n.Related, n.Source, time.Unix(n.Datetime, 0).Format(time.RFC822Z))
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

	return companyNewsList, nil
}
func (cfg *Config) finnhubResCompanyNews(stock string) (*http.Response, error) {
	today := time.Now()
	params := url.Values{}
	params.Add("symbol", stock)
	params.Add("from", today.Format("2006-01-02"))
	params.Add("to", today.Format("2006-01-02"))
	params.Add("token", cfg.FinnhubAPIKey)

	url := finnhubCompNewsURL.ResolveReference(&url.URL{RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
