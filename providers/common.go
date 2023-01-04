package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Parses http.Response with io.ReadAll
func parseWebpageIO(link string) ([]byte, error) {
	// Loads the webpage
	res, err := http.Get(link)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	// Panic if operation is not successful
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Gets the json in bytes
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

// Parses http.response with goquery
func parseWebpageGoquery(link string) (*goquery.Document, error) {
	// Gets webpage data
	res, err := http.Get(link)
	if err != nil {
		return &goquery.Document{}, err
	}
	defer res.Body.Close()

	// Panic if operation is not successful
	if res.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	// Parse response with goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	if err != nil {
		return &goquery.Document{}, err
	}
	return doc, nil
}
