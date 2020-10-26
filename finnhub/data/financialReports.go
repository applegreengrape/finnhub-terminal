package data

import (
	"encoding/json"
	"fmt"
	"github.com/applegreengrape/finnhub-terminal/finnhub"
	"github.com/applegreengrape/finnhub-terminal/finnhub/client"
	"net/url"
	"strings"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const (
	dbPath2 = "./financialReports.db"
)

// GetFinancialReports ..
func GetFinancialReports() (result [][]string, e error) {
	cfg := finnhub.NewSettingFromConfig()
	os.Remove(dbPath2)
	db2, err := sql.Open("sqlite3", dbPath2)
	if err != nil {
		log.Fatal(err)
		//return nil, err
	}
	defer db2.Close()
	loadAllFinancialReports(cfg, db2)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sqlQuery := fmt.Sprintf("select filedDate, type, concept, label, value, unit from %s order by filedDate DESC;", cfg.Stocks[0])
	rows, err := db2.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var d string
		var t string
		var c string
		var l string
		var v string
		var u string
		err = rows.Scan(&d, &t, &c, &l, &v, &u)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		r := []string{}

		if len(d) > 10 {
			d = d[0:10]
		}
		if len(c) > 50 {
			c = c[0:50]+"..."
		}
		if len(l) > 50 {
			l = l[0:50]+"..."
		}
		r = append(r, d, t, c, l, v, u)
		result = append(result, r)
	}

	return result, nil
}

func loadAllFinancialReports(cfg *finnhub.Config, db *sql.DB) error {
	c := &client.Client{
		APIKey: cfg.APIKey,
	}

	path := "/stock/financials-reported"

	for _, s := range cfg.Stocks {
		err := dbinit2(s, db)
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

		id := int64(0)
		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)
		for _, r := range result["data"].([]interface{}) {
			date := fmt.Sprintf("%v", r.(map[string]interface{})["filedDate"])
			report := r.(map[string]interface{})["report"]
			types := report.(map[string]interface{})
			for _, b := range types["bs"].([]interface{}) {
				err := loopLoadDB(b, id, db, s, "bs", date)
				if err != nil {
					log.Fatal(err)
					return err
				}
				id++
			}
			for _, cf := range types["cf"].([]interface{}) {
				err := loopLoadDB(cf, id, db, s, "cf", date)
				if err != nil {
					log.Fatal(err)
					return err
				}
				id++
			}
			for _, ic := range types["ic"].([]interface{}) {
				err := loopLoadDB(ic, id, db, s, "ic", date)
				if err != nil {
					log.Fatal(err)
					return err
				}
				id++
			}

		}

		id = int64(0)
	}

	return nil
}

func loopLoadDB(d interface{}, id int64, db *sql.DB, s, t, date string) error {
	concept := ""
	label := ""
	value := ""
	unit := ""
	for k, v := range d.(map[string]interface{}) {
		switch {
		case k == "concept":
			concept = fmt.Sprintf("%v", v)
		case k == "label":
			label = fmt.Sprintf("%v", v)
		case k == "value":
			value = fmt.Sprintf("%v", v)
		case k == "unit":
			unit = fmt.Sprintf("%v", v)
		}
	}
	err := dbinsert2(db, s, date, t, concept, label, value, unit, id)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func dbinit2(s string, db *sql.DB) error {
	sqlStmt := fmt.Sprintf("create table %s(id int not null primary key, filedDate text, type text, concept text, label text, value text, unit text);", s)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}

func dbinsert2(db *sql.DB, s, d, t, c, l, v, u string, i int64) error {
	concept := strings.ReplaceAll(c, "'", "''")
	label := strings.ReplaceAll(l, "'", "''")
	sqlStmt := fmt.Sprintf("insert into %s(id, filedDate, type, concept, label, value, unit) values('%d', '%s', '%s' ,'%s', '%s', '%s', '%s')", s, i, d, t, concept, label, v, u)
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}
