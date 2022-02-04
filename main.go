package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

type DataFull struct {
	Iterator int
	Data Data
}

type Data struct {
	Title string
	Magnet string
	Torrent string
}


func main() {
	(&cli.App{
		Name: "NyaaLoader",
		Usage: "An efficient Nyaa scraper",
		Commands: []*cli.Command{
			{
				Name: "search",
				Aliases: []string{"s"},
				Usage: "Search term",
				Action: func(c *cli.Context) error {
					fmt.Println("Searching:", c.Args().First())
					Search(c.Args().First())
					return nil
				},
			},
		},
	}).Run(os.Args)
	fmt.Println("Ay Yo, we bootin here!")


	//c.Visit("https://nyaa.si/")

}

func Search(term string)  {

	var DataList []Data

	c := colly.NewCollector()

	c.OnHTML(".torrent-list tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			data := Data{"", "", ""}
			element.ForEach("td a:not(.comments)", func(ii int, a *colly.HTMLElement) {
				if ii == 1 {
					data.Title = a.Text
					//fmt.Printf("Title: %s\n", a.Text)
				}
				MagLink := a.Attr("href")
				if strings.Contains(MagLink, "magnet:") {
					data.Magnet = MagLink
					//fmt.Printf("Magnet: %s\n", MagLink)
				}
				data.Torrent = ""
				//if strings.Contains(MagLink, ".torrent") {
				//	data.Torrent = MagLink
				//	fmt.Printf("Torrent: %s\n", MagLink)
				//}
			})
			res := append(DataList, data)

			fmt.Printf("Data: %s\n", res)
		})
		//link := e.ForEach("td")

	})



	err := c.Visit("https://nyaa.si/?q=" + term)
	if err != nil {
		return
	}
}
/*
Returns:
Ay Yo, we bootin here!
Magnet magnet:?xt=urn:btih:ced4242cfd62d521f0c8f7527b31b5be809d4e64&dn=%E3%80%90%E9%80%A3%E7%B6%9A%E3%83%86%E3%83%AC%E3%83%93%E5%B0%8F%E8%AA%AC%E3%80%91%E3%82%AB%E3%83%BC%E3%83%8D%E3%83%BC%E3%82%B7%E3%83%A7%E3%83%B3%20%E7%AC%AC13%E9%80%B1%E3%80%8C%E7%94%9F%E3%81%8D%E3%82%8B%E3%80%8D%20-%20%5BBS12%5D%5B1920x1080p.EP073-075.hevc.mkv%5D%5B%E7%84%A1%5D&tr=http%3A%2F%2Fnyaa.tracker.wf%3A7777%2Fannounce&tr=udp%3A%2F%2Fopen.stealth.si%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Fexodus.desync.com%3A6969%2Fannounce&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce
... More Magnet links
*/
