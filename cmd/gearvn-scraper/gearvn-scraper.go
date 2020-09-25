package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

const shopURL = "https://gearvn.com"

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("gearvn.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "products") {
			visited, err := c.HasVisited(link)
			if err != nil {
				e.Request.Abort()
			}
			if !visited {
				fmt.Printf("Link found: %q -> %s\n", e.Text, link)
				c.Visit(e.Request.AbsoluteURL(link))
			}

		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(shopURL)
}
