package providers

import (
	"fmt"
	"net/http"
	"strings"
)

func UnsplashWallpaper(height, width int, crop bool) (string, error) {
	// Asks source.unsplash.com for random image
	URL := "https://source.unsplash.com/random"

	// If height or width are 0, it will return a 16:9 image as fallback
	if height <= 0 || width <= 0 {
		URL += "/16x9"
	} else {
		URL += fmt.Sprintf("/%dx%d", width, height)
	}

	// Get image url (via redirects)
	res, err := http.Head(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Checking that the status code is 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// getting image url from the final redirect
	imageURL := res.Request.URL.String()

	// extracting ixid
	_, ixid, _ := strings.Cut(imageURL, "&ixid=")
	ixid, _, _ = strings.Cut(ixid, "&ixlib")
	ixid = "?ixid=" + ixid

	// Compose final url, stripping all queries and readding ixid and format
	finalURL, _, _ := strings.Cut(imageURL, "?")
	finalURL += ixid + "&fm=png"

	if height > 0 || width > 0 {
		finalURL += fmt.Sprintf("&h=%d&w=%d", height, width)
	}

	if crop {
		finalURL += "&crop=entropy&fit=crop"
	} else {
		finalURL += "&fit=max"
	}

	return finalURL, nil
}
