package yahoo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
)

// GetHistPrice ..
func GetHistPrice() (prices []float64) {
	cfg := NewSettingFromConfig()
	hist, err := cfg.getHistPrice(cfg.Stocks[0])
	if err != nil {
		panic(err)
	}

	for _, p := range hist {
		prices = append(prices, round(p, .5, 2))
	}

	return prices
}

/* -------------------- Unexported Functions -------------------- */

func (cfg *Config) getHistPrice(s string) ([]float64, error) {
	h := []float64{}

	resp, err := cfg.yahooHistRequest(s)
	if err != nil {
		return nil, err
	}

	var res YahooResp
	json.NewDecoder(resp.Body).Decode(&res)
	for _, c := range res.Chart.Result[0].Indicators.Quote[0].Close {
		h = append(h, c)
	}

	return h, nil
}

func (cfg *Config) yahooHistRequest(symbol string) (*http.Response, error) {
	params := url.Values{}
	params.Add("interval", "1m")
	params.Add("period", "1d")

	url := yahooURL.ResolveReference(&url.URL{Path: symbol, RawQuery: params.Encode()})

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	//rand.Seed(86)
	//n := rand.Intn(len(proxyServers))
	//proxyURL, err := url.Parse(proxyServers[n])
	///httpClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
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

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
