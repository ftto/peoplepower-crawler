package main

import (
	"time"

	"github.com/gocolly/colly"
)

type PersonName struct {
	Ko    string `JSON:"ko"`
	HanJa string `JSON:"hanJa"`
}

type SangIm struct {
	Link string `JSON:"link"`
	Text string `JSON:"text"`
}

type Person struct {
	ID       string     `JSON:"id"`
	Name     PersonName `JSON:"name"`
	Party    string     `JSON:"party"`
	Precinct string     `JSON:"precinct"`
	SangIm   SangIm     `JSON:"sangIm"`
}

var DefaultCrawler *colly.Collector
var People map[string]*Person

func init() {
	DefaultCrawler = colly.NewCollector(
		colly.Async(true),
	)

	DefaultCrawler.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
}

func main() {
	PaginationCrawler.Visit("http://watch.peoplepower21.org/?act=&mid=AssemblyMembers&vid=&mode=search")
	PaginationCrawler.Wait()
	PeopleCrawler.Wait()
	PersonCrawler.Wait()

	// for v := range People {
	// 	fmt.Println(People[v])
	// }
}
