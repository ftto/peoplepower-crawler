package main

import (
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

var PersonCrawler *colly.Collector

func init() {
	PersonCrawler = DefaultCrawler.Clone()

	PersonCrawler.OnRequest(func(r *colly.Request) {
		m, _ := url.ParseQuery(r.URL.RawQuery)
		id := m["member_seq"][0]
		r.Ctx.Put("id", id)
	})

	// 1. Name
	PersonCrawler.OnHTML(".panel-default > .panel-body > h1", func(e *colly.HTMLElement) {
		names := strings.Split(strings.TrimSpace(e.Text), "  ")
		e.Request.Ctx.Put("nameKo", names[0])
		e.Request.Ctx.Put("nameHanJa", names[1])
	})

	// 2. profile image
	PersonCrawler.OnHTML(".panel-default > .panel-body > .row > .col-md-3 > img", func(e *colly.HTMLElement) {
		profileImg := e.Attr("src")
		e.Request.Ctx.Put("profileImg", profileImg)
	})

	// 3. Personal Information
	PersonCrawler.OnHTML(".panel-body", func(e *colly.HTMLElement) {
		e.ForEach("table.table-user-information tbody tr", func(index int, tr *colly.HTMLElement) {
			switch index {
			case 0:
				// 정당
				party := strings.Trim(strings.Replace(tr.ChildText("td:nth-of-type(2)"), tr.ChildText("td:nth-of-type(2) > span"), "", 1), " ")
				e.Request.Ctx.Put("party", party)
			case 1:
				// 선거구
				precinct := tr.ChildText("td:nth-of-type(2) > a")
				e.Request.Ctx.Put("precinct", precinct)
			case 2:
				// 당선횟수
			case 3:
				// 소속위원회
				sangImLink := tr.ChildAttr("td:nth-of-type(2) > a", "href")
				sangImText := tr.ChildText("td:nth-of-type(2) > a")
				e.Request.Ctx.Put("sangimLink", sangImLink)
				e.Request.Ctx.Put("sangimText", sangImText)
			case 4:
				// 학력
				education := strings.Replace(tr.ChildText("td:nth-of-type(2)"), "<br>", ";", -1)
				e.Request.Ctx.Put("education", education)
			case 5:
				// 경력
				career := strings.Replace(tr.ChildText("td:nth-of-type(2)"), "<br>", ";", -1)
				e.Request.Ctx.Put("career", career)
			case 7:
				// 경력
				email := tr.ChildText("td:nth-of-type(2) > a")
				e.Request.Ctx.Put("email", email)
			default:
			}
		})
	})

	PersonCrawler.OnScraped(func(r *colly.Response) {
		person := &Person{
			ID: r.Ctx.Get("id"),
			Name: PersonName{
				Ko:    r.Ctx.Get("nameKo"),
				HanJa: r.Ctx.Get("nameHanJa"),
			},
			Party:    r.Ctx.Get("party"),
			Precinct: r.Ctx.Get("precinct"),
			SangIm: SangIm{
				Link: r.Ctx.Get("sangImLink"),
				Text: r.Ctx.Get("sangImText"),
			},
		}

		People[person.ID] = person
	})
}
