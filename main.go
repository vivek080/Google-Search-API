package main

import (
	"bufio"
	"fmt"
	"os"

	googlescraper "./scraper"
)

func main() {

	fmt.Println("Enter the Keyword you want to search: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keyword := scanner.Text()

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the Keyword you want to search: ")
	// keyword, _ := reader.ReadString('\n')
	// keyword = strings.Replace(keyword, "\n", "", -1)

	fmt.Printf("\nplease wait while it is searching for the keyword '%s' \n", keyword)
	fmt.Println()

	// passing the keyword to my GoogleScrape handler
	res, err := googlescraper.GoogleScrape(keyword) // getting the results in array form
	if err != nil {
		fmt.Printf("Did not find any results, please check your entered Keyword \nreturned error '%s'\n", err)
	}

	for count, item := range res {
		if count < 5 {
			fmt.Println(item)
		} else {
			break
		}
	}
}
