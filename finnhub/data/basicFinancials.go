package data

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	dbPath = "./basicFinancials.db"
)

//GetBasicFinancials ..
func GetBasicFinancials() (result [][]string, e error) {
	cfg := finnhub.NewSettingFromConfig()
	os.Remove(dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	err = loadAllBasicFinancials(cfg, db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sel := ""
	joins := ""
	for _, s := range cfg.Stocks {
		if s == cfg.Stocks[len(cfg.Stocks)-1] {
			sel += fmt.Sprintf("%s.%s ", s, s)
		} else {
			sel += fmt.Sprintf("%s.%s, ", s, s)
		}
		if s != cfg.Stocks[0] {
			joins += fmt.Sprintf("left join %s on %s.metric = %s.metric ", s, cfg.Stocks[0], s)
		}
	}

	sqlQuery := fmt.Sprintf("select %s.metric, %s from %s %s order by %s.metric;", cfg.Stocks[0], sel, cfg.Stocks[0], joins, cfg.Stocks[0])
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		row := make([]string, len(cols))

		for i, raw := range rawResult {
			if raw == nil {
				row[i] = "\\N"
			} else {
				row[i] = string(raw)
			}
		}

		result = append(result, row)
	}

	return result, nil
}

func loadAllBasicFinancials(cfg *finnhub.Config, db *sql.DB) error {
	c := &client.Client{
		APIKey: cfg.APIKey,
	}

	path := "/stock/metric"
	for _, s := range cfg.Stocks {
		err := dbinit(s, db)
		if err != nil {
			log.Fatal(err)
			return err
		}
		p := url.Values{}
		p.Add("symbol", s)
		res, err := c.FinnhubClient(finnhub.Version+path, p)
		if err != nil {
			log.Fatal(err)
			return err
		}

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)
		m := result["metric"].(map[string]interface{})

		for k, v := range m {
			err := dbinsert(db, s, k, fmt.Sprintf("%v", v))
			if err != nil {
				log.Fatal(err)
				return err
			}
		}

	}

	return nil
}

func dbinit(s string, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("create table %s(metric text not null primary key, %s text);", s, s)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func dbinsert(db *sql.DB, s, m, v string) error {
	sqlStmt := fmt.Sprintf("insert into %s(metric, %s) values('%s', '%s')", s, s, m, v)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
