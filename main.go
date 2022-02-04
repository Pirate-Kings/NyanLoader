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
/*
Returns:
Ay Yo, we bootin here!
Magnet magnet:?xt=urn:btih:ced4242cfd62d521f0c8f7527b31b5be809d4e64&dn=%E3%80%90%E9%80%A3%E7%B6%9A%E3%83%86%E3%83%AC%E3%83%93%E5%B0%8F%E8%AA%AC%E3%80%91%E3%82%AB%E3%83%BC%E3%83%8D%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%20%E7%AC%AC13%E9%80%B1%E3%80%8C%E7%94%9F%E3%81%8D%E3%82%8B%E3%80%8D%20-%20%5BBS12%5D%5B1920x1080p.EP073-075.hevc.mkv%5D%5B%E7%84%A1%5D&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce
... More Magnet links
*/
