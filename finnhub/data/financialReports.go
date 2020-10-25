package data

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"log"
	"net/url"
)

func loadAllFinancialReports(cfg *finnhub.Config) error {
	c := &client.Client{
		APIKey: cfg.APIKey,
	}

	path := "/stock/financials-reported"
	//err := dbinit(s, db)
	//if err != nil {
	//	log.Fatal(err)
	//	return err
	//}
	p := url.Values{}
	p.Add("symbol", cfg.Stocks[0])
	res, err := c.FinnhubClient(finnhub.Version+path, p)
	if err != nil {
		log.Fatal(err)
		return err
	}

	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	//data := result["data"].(map[string]interface{})
	//fmt.Println(data)
	rows:= [][]string{}
	row := []string{}
	for _, r := range result["data"].([]interface{}) {
		_report := r.(map[string]interface{})["report"]
		for _, bs := range _report.(map[string]interface{}) {
			for _, c := range bs.([]interface{}) {
				for k, v := range c.(map[string]interface{}) {
					fmt.Println(k, ":", v)
					//if k == ""
					row = append(row, fmt.Sprintf("%v", v))
				}
				rows = append(rows, row)
				row = []string{}
				fmt.Println("")
			}
		}
	}

	fmt.Println(rows[0])

	return nil
}
