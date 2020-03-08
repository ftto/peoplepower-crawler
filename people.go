package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var PeopleCrawler *colly.Collector

func init() {
	People = make(map[string]*Person)
	PeopleCrawler = DefaultCrawler.Clone()

	// Fetch List of People
	PeopleCrawler.OnRequest(func(r *colly.Request) {
		fmt.Println("On Request")
	})

	PeopleCrawler.OnHTML(".col-md-8 > .col-xs-6 a[href]", func(e *colly.HTMLElement) {
		PersonCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org%s", e.Attr("href")))
	})
}
