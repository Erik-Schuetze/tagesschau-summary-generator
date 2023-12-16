package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Article struct {
	Headline string
	Content  string
	Link     string
	Tags     []string
}

func main() {
	articles := []Article{}

	links := getNews()

	for _, link := range links {
		articles = append(articles, getArticle(link))
	}

	for _, e := range articles {
		fmt.Println(e.Link)
		fmt.Println(e.Headline)
		fmt.Println(e.Tags)
		fmt.Println(e.Content)
		fmt.Println(" ")
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

func getArticle(link string) Article {
	article := new(Article)
	article.Link = link

	c := colly.NewCollector()

	// get headline
	c.OnHTML("div>h1>span", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("class"), "headline") {
			article.Headline = e.Text
		}
	})

	// fuse texts together
	content := ""
	c.OnHTML("article>p", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("class"), "textabsatz") {
			content = fmt.Sprintf("%s%s ", content, strings.TrimSpace(e.Text))
		}
		article.Content = content
	})

	// get tags
	tags := []string{}
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("class"), "tag-btn") {
			tags = append(tags, e.Text)
		}
		article.Tags = tags
	})

	c.Visit(link)

	return *article
}
