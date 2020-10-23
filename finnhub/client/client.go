package client

import (
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"net/http"
	"net/url"
)

// base url and version for finnhub api
var (
	baseURL = &url.URL{Scheme: "https", Host: finnhub.BaseURL}
)

// Client ..
type Client struct {
	APIKey string `json:"apiKey"`
}

// FinnhubClient Client
func (c *Client) FinnhubClient(path string, params url.Values) (*http.Response, error) {
	params.Add("token", c.APIKey)

	url := baseURL.ResolveReference(&url.URL{Path: path, RawQuery: params.Encode()})
	req, err := http.NewRequest("GET", url.String(), nil)
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
