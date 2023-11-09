package providers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/tidwall/gjson"
)

func BingWallpaper(market, resolution string, quality, height, width, crop int) (string, error) {
	// Generates the bing api url, randomizing the index and market
	bingURL := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=1&mkt=%s", rand.Intn(7), market)

	// Get Json
	res, err := http.Get(bingURL)
	if err != nil {
		return "", err
	}

	// ignoring unhandled error here, shouldn't be a problem and don't want to over complicate too much
	defer res.Body.Close()

	// Error "handling"
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Gets the json in bytes
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Parse urlbase
	imageURLBase := gjson.GetBytes(content, "images.0.urlbase").String()

	// Compose final url
	finalURL := fmt.Sprintf("http://bing.com%s_%s.jpg&qlt=%d&p=0&pid=hp", imageURLBase, resolution, quality)

	// add height and width parameters if higher than 0
	if height > 0 {
		finalURL += fmt.Sprintf("&h=%d", height)
	}
	if width > 0 {
		finalURL += fmt.Sprintf("&w=%d", width)
	}

	// add crop
	switch crop {
	case 1:
		finalURL += "&c=4" // blind ratio

	case 2:
		finalURL += "&c=7" // smart ratio
	}

	return finalURL, nil
}
