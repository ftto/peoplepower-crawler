package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var PaginationCrawler *colly.Collector

func init() {
	PaginationCrawler = DefaultCrawler.Clone()

	PaginationCrawler.OnHTML("ul.pagination", func(e *colly.HTMLElement) {
		// pages := e.DOM.Children().Length()

		for i := 0; i < 1; i++ {
			PeopleCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search&page=%d", (i + 1)))
		}
	})
}
