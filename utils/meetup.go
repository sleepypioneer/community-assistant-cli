package utils

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type Meetup struct {
	Title     string
	Time      string
	Attendees float64
}

type Scraper struct {
	collector *colly.Collector
}

// scrapes a meetup.com page to get the meetup title, date and number of attendees
func ScrapeMeetupPage(eventID string) Meetup {

	c := colly.NewCollector(
		colly.AllowedDomains("www.meetup.com"),
	)

	scraper := Scraper{
		collector: c,
	}

	title, time, attendees := scraper.ScrapeTitle(eventID)

	fmt.Printf("Title: %s\n Time: %s\n Attendees %f\n", title, time, attendees)

	return Meetup{
		Title:     title,
		Time:      time,
		Attendees: attendees,
	}
}

func (s Scraper) ScrapeTitle(eventID string) (string, string, float64) {
	var title, time string
	var attendees float64

	s.collector.OnHTML("#main", func(e *colly.HTMLElement) {
		dom := e.DOM

		title = dom.Find("h1").Text()
	})

	s.collector.OnHTML("html", func(e *colly.HTMLElement) {
		e.DOM.Find("script").Each(func(i int, s *goquery.Selection) {
			buffer := []byte(s.Text())
			var bufferSingleMap map[string]interface{}
			json.Unmarshal(buffer, &bufferSingleMap)

			if bufferSingleMap["name"] == title {
				time = bufferSingleMap["startDate"].(string)
			}

			if bufferSingleMap["@type"] == "RsvpAction" {
				attendees = bufferSingleMap["additionalNumberOfGuests"].(float64)
			}
		})

	})

	s.collector.Visit(fmt.Sprintf("https://www.meetup.com/pyladies-berlin/events/%s/", eventID))

	return title, time, attendees
}
