package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

func NasaWallpaper(apiKey string) (string, error) {
	URL := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&count=1&hd=true", apiKey)

	// Get JSON
	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Checking for rate limiting
	if res.Header.Get("X-RateLimit-Remaining") == "0" {
		return "", fmt.Errorf("you're getting rate limited")
	}

	// Checking that the status code is 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Gets the response in bytes to decode it with gjson
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// parse the json and return the image
	imageUrl := gjson.GetBytes(content, "0.hdurl").String()
	return imageUrl, nil
}
