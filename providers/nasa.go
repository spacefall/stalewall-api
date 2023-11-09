package providers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

// NasaWallpaper - Gets a wallpaper from Nasa Apod
func NasaWallpaper(apiKey string) (string, error) {
	apiURL := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&count=1&hd=true", apiKey)

	var imageUrl string

	for gotImg := false; !gotImg; gotImg = imageUrl != "" {
		// Querying the api
		res, err := http.Get(apiURL)
		if err != nil {
			return "", err
		}

		// ignoring error, but should check the defer in loop thing
		defer res.Body.Close()

		// handling errors
		// general errors
		if res.StatusCode != http.StatusOK {
			return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// rate limiting
		if res.Header.Get("X-RateLimit-Remaining") == "0" {
			return "", fmt.Errorf("you have exceeded your rate limit")
		}

		// getting json from response
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}

		// Parse JSON
		imageUrl = gjson.GetBytes(body, "0.hdurl").String()

	}

	//println(gjson.ParseBytes(body).String())

	return imageUrl, nil

}
