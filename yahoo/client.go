package yahoo

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
)

// GetPrice ..
func (client *Client) GetPrice() ([]StockPrice, error) {
	quotes := []StockPrice{}

	for _, s := range client.symbols {
		resp, err := client.yahooRequest(s)
		if err != nil {
			return quotes, err
		}

		var res YahooResp
		var quote StockPrice
		quote.stock = s
		json.NewDecoder(resp.Body).Decode(&res)
		quote.price = res.Chart.Result[0].Meta.RegularMarketPrice
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

/* -------------------- Unexported Functions -------------------- */

var (
	yahooURL = &url.URL{Scheme: "https", Host: "query1.finance.yahoo.com", Path: "/v8/finance/chart/"}
)

var proxyServers = []string{
	"192.168.1.1:1080",
	"103.52.211.186:1080",
	"176.9.75.42:1080",
	"207.154.231.213:1080",
	"138.68.161.14:1080",
	"193.34.161.137:44436",
	"95.79.112.74:3629",
	"138.197.157.32:8080",
	"181.129.51.147:47562",
	"90.181.150.211:4145",
	"",
}

func (client *Client) yahooRequest(symbol string) (*http.Response, error) {
	params := url.Values{}
	params.Add("interval", "1d")
	params.Add("period", "1d")

	url := yahooURL.ResolveReference(&url.URL{Path: symbol, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	rand.Seed(86)
	n := rand.Intn(len(proxyServers))
	proxyURL, err := url.Parse(proxyServers[n])
	httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
