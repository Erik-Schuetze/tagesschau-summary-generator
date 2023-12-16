package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	articles := getNews()

	for _, element := range articles {
		//fmt.Println(element)
		content := getArticleContent(element)
		fmt.Printf("\n%s:\n%s", element, content)
	}

}

func getNews() []string {
	articles := []string{}

	// https://www.tagesschau.de/archiv/allemeldungen?datum=2023-12-14
	timestamp := time.Now().Format("2006-01-02")
	startUrl := fmt.Sprintf("https://www.tagesschau.de/archiv/allemeldungen?datum=%s", timestamp)

	c := colly.NewCollector()

	c.OnHTML("ul>div>li>a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		articles = append(articles, e.Request.AbsoluteURL(link))
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.Visit(startUrl)

	return articles
}

func getArticleContent(link string) string {
	articleContent := ""

	c := colly.NewCollector()

	c.OnHTML("article>p", func(e *colly.HTMLElement) {
		articleContent = fmt.Sprintf("%s%s ", articleContent, strings.TrimSpace(e.Text))
	})

	c.Visit(link)

	return articleContent
}
