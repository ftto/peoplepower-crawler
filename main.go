package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	PaginationCrawler := colly.NewCollector()
	PeopleCrawler := colly.NewCollector(
		colly.Async(true),
	)
	PeopleCrawler.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	PersonCrawler := PeopleCrawler.Clone()

	// Fetch Pages
	PaginationCrawler.OnHTML("ul.pagination", func(e *colly.HTMLElement) {
		// pages := e.DOM.Children().Length()

		for i := 0; i < 1; i++ {
			fmt.Println(i)
			PeopleCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search&page=%d", (i + 1)))
		}
	})

	// Fetch List of People
	PeopleCrawler.OnRequest(func(r *colly.Request) {
		fmt.Println("On Request")
	})

	PeopleCrawler.OnHTML(".col-md-8 > .col-xs-6 a[href]", func(e *colly.HTMLElement) {
		PersonCrawler.Visit(fmt.Sprintf("http://watch.peoplepower21.org%s", e.Attr("href")))
	})

	// Fetch Every Person
	PersonCrawler.OnHTML(".panel-default > .panel-body > h1", func(e *colly.HTMLElement) {
		// names := strings.Split(strings.TrimSpace(e.Text), "  ")
		// ko := names[0]
		// hanja := names[1]
	})

	PaginationCrawler.Visit("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search")
	PeopleCrawler.Wait()
	PersonCrawler.Wait()
}
