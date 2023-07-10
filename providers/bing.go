package providers

import (
	"fmt"
	"math/rand"

	"github.com/tidwall/gjson"
)

func BingWallpaper(market, resolution string, quality, height, width, crop int) (string, error) {
	// Generates the bing api url, randomizing the index and market
	bingURL := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=1&mkt=%s", rand.Intn(7), market)

	// Get Json
	content, err := parseWebpageIO(bingURL)
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
