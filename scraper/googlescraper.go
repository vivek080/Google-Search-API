package googlescraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GoogleResult struct {
	ResultURL string
}

func buildGoogleURL(searchTerm string) string { //building the google Search url which will be used to query for the given keyword
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	languageCode := "en"
	googleBase := "https://www.google.com/search?q="
	return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
}

func googleRequest(searchURL string) (*http.Response, error) { //requesting the web browser the open the link which is build for querying the keyqord
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func googleResultParser(response *http.Response) ([]GoogleResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []GoogleResult{}
	sel := doc.Find("div.g") //extracting all the individual link block to sel variable
	// fmt.Println(sel.Length())
	// fmt.Println(response)
	// fmt.Println(sel.Nodes[0])
	for i := range sel.Nodes {
		item := sel.Eq(i)               //selecting each link block to extract the href link
		linkTag := item.Find("a")       //finding the anchor tag from the block
		link, _ := linkTag.Attr("href") //extracting the href link from the anchor block
		if link != "" && link != "#" {
			result := GoogleResult{
				link,
			}
			results = append(results, result)
		}
	}
	return results, err
}

func GoogleScrape(searchTerm string) ([]GoogleResult, error) {
	googleURL := buildGoogleURL(searchTerm) //building the search URL using the Search keyword
	res, err := googleRequest(googleURL)    //accesing the Link in browser client using the above build search URL
	if err != nil {
		return nil, err
	}
	scrapes, err := googleResultParser(res) //Scraping the url from the browser web page using the "github.com/PuerkitoBio/goquery" functions
	if err != nil {
		return nil, err
	}
	return scrapes, nil
}
