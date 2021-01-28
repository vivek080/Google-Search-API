package main

import (
	"fmt"

	googlescraper "./scraper"
)

func main() {
	keyword := "vivek hiremath"
	// passing the keyword to my GoogleScrape handler
	res, _ := googlescraper.GoogleScrape(keyword) // getting the results in array form
	count := 0
	for _, item := range res {
		if count < 5 {
			fmt.Println(item)
		} else {
			break
		}
		count++
	}
}
