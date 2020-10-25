package data

import (
	"github.com/applegreengrape/finnhub-terminal/finnhub"

	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// ExportAllBasicFinancials ..
func ExportAllBasicFinancials() error {
	cfg := finnhub.NewSettingFromConfig()

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	for _, s := range cfg.Stocks {
		os.Remove(fmt.Sprintf("%s.basicFinancials.csv", s))
		f, err := os.Create(fmt.Sprintf("%s.basicFinancials.csv", s))
		w := csv.NewWriter(f)
		defer w.Flush()
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = w.Write([]string{"metric", "value"})
		if err != nil {
			log.Fatal(err)
			return err
		}

		sqlQuery := fmt.Sprintf("SELECT * from %s ORDER by %s.metric;", s, s)
		rows, err := db.Query(sqlQuery)
		if err != nil {
			log.Fatal(err)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			var m string
			var v string
			err = rows.Scan(&m, &v)
			if err != nil {
				log.Fatal(err)
				return err
			}
			err := w.Write([]string{m, v})
			if err != nil {
				log.Fatal(err)
				return err
			}
		}

	}
	return nil
}

// ExportAllFinancialReports ..
func ExportAllFinancialReports() error {
	cfg := finnhub.NewSettingFromConfig()

	db2, err := sql.Open("sqlite3", dbPath2)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db2.Close()

	for _, s := range cfg.Stocks {
		os.Remove(fmt.Sprintf("%s.financialReports.csv", s))
		f, err := os.Create(fmt.Sprintf("%s.financialReports.csv", s))
		w := csv.NewWriter(f)
		defer w.Flush()
		if err != nil {
			log.Fatal(err)
			return err
		}
		err = w.Write([]string{"filedDate", "type", "concept", "label", "value", "unit"})
		if err != nil {
			log.Fatal(err)
			return err
		}

		sqlQuery2 := fmt.Sprintf("select filedDate, type, concept, label, value, unit from %s order by filedDate DESC;", s)
		rows, err := db2.Query(sqlQuery2)
		if err != nil {
			log.Fatal(err)
			return err
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
				return err
			}
			err := w.Write([]string{d, t, c, l, v, u})
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}
