package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Product struct {
	Name   string
	Rating float64
	Price  float64
}

var ProductList []Product

func compareRating(i, j int) bool {
	return ProductList[i].Rating > ProductList[j].Rating
}

func ScrapeData(keyword string) {

	c := colly.NewCollector((colly.AllowedDomains("www.amazon.in")))

	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Link of the page:", r.URL, "\n \n")
	})

	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-small.a-spacing-top-small", func(_ int, h *colly.HTMLElement) {
			name := h.ChildText("span.a-color-base.a-text-normal")
			stars := h.ChildText("span.a-icon-alt")
			price := h.ChildText("span.a-price-whole")

			var product Product

			if name != "" {
				s := strings.Split(stars, " ")
				priceInFloat, _ := strconv.ParseFloat(price, 64)
				ratingInFloat, _ := strconv.ParseFloat(s[0], 64)

				product.Name = name
				product.Rating = ratingInFloat
				product.Price = priceInFloat
				ProductList = append(ProductList, product)
			}
		})
	})

	c.Visit("https://www.amazon.in/s?k=" + keyword)

	sort.Slice(ProductList, compareRating)

	fmt.Print("Top 5 Products for the given keyword: \n \n")
	for i := 0; i < 5; i++ {

		fmt.Println("Ranking: #", i+1)
		fmt.Println("Name:", ProductList[i].Name)
		fmt.Println("Rating:", ProductList[i].Rating)
		// if !reflect.ValueOf(ProductList[i].Price).IsZero() {
		// 	fmt.Println("Price:", ProductList[i].Price)
		// } else {
		// 	fmt.Println("Price: NA")
		// }
		fmt.Println("")
	}
}
