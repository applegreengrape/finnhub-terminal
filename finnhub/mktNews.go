package finnhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// MarketNews ..
func MarketNews() []string {
	cfg := NewSettingFromConfig()

	news, err := cfg.getMktNews()
	if err != nil {
		//
	}

	return news
}

/* -------------------- Unexported Functions -------------------- */

var (
	finnhubMktNewsURL = &url.URL{Scheme: "https", Host: "finnhub.io", Path: "/api/v1/news"}
)

func (cfg *Config) getMktNews() ([]string, error) {

	mktNewsList := []string{}
	resp, err := cfg.finnhubMktNews()
	if err != nil {
		return nil, err
	}

	var mktnews News
	json.NewDecoder(resp.Body).Decode(&mktnews)

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

	return mktNewsList, nil
}

func (cfg *Config) finnhubMktNews() (*http.Response, error) {
	params := url.Values{}
	params.Add("category", "general")
	params.Add("token", cfg.FinnhubAPIKey)

	url := finnhubMktNewsURL.ResolveReference(&url.URL{RawQuery: params.Encode()})

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
