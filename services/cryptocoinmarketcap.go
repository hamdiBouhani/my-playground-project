package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func give_current_price_of_bitcoin() error {

	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)", "Circulating Supply (USD)", "Volume (24h)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".cmc-table__column-name"),
			e.ChildText(".cmc-table__cell--sort-by__symbol"),
			e.ChildText(".cmc-table__cell--sort-by__market-cap"),
			e.ChildText(".cmc-table__cell--sort-by__price"),
			e.ChildText(".cmc-table__cell--sort-by__circulating-supply"),
			e.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
		})
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err = c.Visit("https://coinmarketcap.com/all/views/all/")

	if err != nil {
		fmt.Println(err)
		return err
	}

	log.Printf("Scraping finished, check file %q for results\n", fName)
	return nil
}
