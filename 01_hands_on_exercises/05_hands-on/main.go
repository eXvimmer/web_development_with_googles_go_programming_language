package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

var tmpl *template.Template

// returns a nicely formatted UTC string representation of a time.Time object
func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006")
}

type MarketData struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

func init() {
	var functions = template.FuncMap{
		"humanDate": humanDate,
	}
	tmpl = template.Must(template.New("index.tmpl.html").Funcs(functions).ParseFiles("./index.tmpl.html"))
}

func main() {
	records, err := parseMarketDataFromCSV("./table.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatal(err)
	}
}

func parseMarketDataFromCSV(filename string) ([]*MarketData, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []*MarketData{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)

	// skip the first line
	if _, err := r.Read(); err != nil {
		return []*MarketData{}, err
	}

	var records []*MarketData

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []*MarketData{}, err
		}
		t, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return []*MarketData{}, err
		}
		o, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		h, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		l, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		c, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		v, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		ac, err := strconv.ParseFloat(record[6], 64)
		if err != nil {
			return []*MarketData{}, err
		}
		data := &MarketData{
			Date:     t,
			Open:     o,
			High:     h,
			Low:      l,
			Close:    c,
			Volume:   v,
			AdjClose: ac,
		}

		records = append(records, data)
	}

	return records, nil
}
