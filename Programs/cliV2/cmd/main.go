package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "soccer_data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("espn.com", "www.espn.com"), // Assuming ESPN domain, adjust as necessary
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"),
	)

	c.OnHTML("tbody.Table__TBODY", func(e *colly.HTMLElement) {
		e.ForEach("tr.Table__TR--sm", func(_ int, el *colly.HTMLElement) {
			team1 := el.ChildText("span.Table__Team.away a.AnchorLink")
			team2 := el.ChildText("span.Table__Team a.AnchorLink:last-child")
			gameStatus := el.ChildText("span.gameNote")
			time := el.ChildText("td.date__col a.AnchorLink")
			venue := el.ChildText("td.venue__col div")
			gameStatusV2 := el.ChildText("td.teams__col a.AnchorLink")

			if gameStatusV2 == "FT" {
				gameStatus = "Full Time"

				if time == "" {
					time = gameStatusV2
				}
			}

			details := []string{team1, team2, gameStatus, venue, time}

			if err := writer.Write(details); err != nil {
				log.Fatalf("Could not write to CSV, err: %q", err)
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.Visit("https://www.espn.com/soccer/schedule") // Modify to the actual URL you are targeting

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}

func displayLess() {

}
