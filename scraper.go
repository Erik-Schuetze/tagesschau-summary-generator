package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	articles := getNews()

	for _, element := range articles {
		fmt.Println(element)
	}

}

func getNews() []string {
	articles := []string{}

	// https://www.tagesschau.de/archiv/allemeldungen?datum=2023-12-14
	timestamp := time.Now().Format("2006-01-02")
	startUrl := fmt.Sprintf("https://www.tagesschau.de/archiv/allemeldungen?datum=%s", timestamp)

	c := colly.NewCollector(
	/*colly.URLFilters(
		regexp.MustCompile(startUrl),
		regexp.MustCompile(`https:\/\/www\.tagesschau\.de.*`),
	),*/
	)

	c.OnHTML("ul>div>li>a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		articles = append(articles, e.Request.AbsoluteURL(link))
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(startUrl)

	return articles
}
