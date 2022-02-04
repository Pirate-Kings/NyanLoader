package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	fmt.Println("Ay Yo, we bootin here!")
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "magnet:") {
			fmt.Printf("Magnet %s\n", link)
		}
	})

	c.Visit("https://nyaa.si/")
}
