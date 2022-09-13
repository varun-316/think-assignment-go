package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func scrapeData() {
	c := colly.NewCollector((colly.AllowedDomains("www.amazon.in")))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		// fmt.Print(h)
		h.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(_ int, h *colly.HTMLElement) {
			name := h.ChildText("span.a-color-base.a-text-normal")
		// 	// var stars string
		// 	// stars = h.ChildText("span.a-icon-alt")
		// 	// var price string
		// 	// price = h.ChildText("span.a-price-whole")
			
			fmt.Println("Product name:", name + "\n");
		})
	})

	c.Visit("https://www.amazon.in/s?k=keyboard")
}