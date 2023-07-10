package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Parses http.Response with io.ReadAll
// Used for apis that retun a json e.g. bing
func parseWebpageIO(link string) ([]byte, error) {
	// Loads the webpage
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	/* 	client := &http.Client{}
	   	req, err := http.NewRequest("GET", link, nil)
	   	if err != nil {
	   		return nil, err
	   	}
	   	req.Header.Set("User-Agent", "WindowsShellClient/0")
	   	res, err := client.Do(req)
	   	if err != nil {
	   		return nil, err
	   	} */
	defer res.Body.Close()

	// return an error if operation is not successful
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Gets the json in bytes
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Parses http.response with goquery
// Used for sources without api (so we scrape the site)
func parseWebpageGoquery(link string) (*goquery.Document, error) {
	// Gets webpage data
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// return an error if operation is not successful
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Parse response with goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
