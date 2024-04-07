package providers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/tidwall/gjson"
)

var bMarkets = [27]string{"es-AR", "en-AU", "de-AT", "nl-BE", "fr-BE", "pt-BR", "en-CA", "fr-CA", "da-DK", "fi-FI", "fr-FR", "de-DE", "zh-HK", "en-IN", "en-ID", "it-IT", "ja-JP", "ko-KR", "zh-CN", "pl-PL", "ru-RU", "es-ES", "sv-SE", "tr-TR", "en-GB", "en-US", "es-US"}

func BingWallpaper(height, width int, crop bool) (string, error) {
	// Gets a random index (day of the week) and market and stitches it to the api url
	URL := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=1&mkt=%s", rand.Intn(7), bMarkets[rand.Intn(len(bMarkets))])

	// Get Json
	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Checking that the status code is 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Gets the response in bytes to decode it with gjson
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	imageURLBase := gjson.GetBytes(content, "images.0.urlbase").String()

	// Compose final url, last 2 options are file resolution and quality
	finalURL := fmt.Sprintf("https://bing.com%s_%s.jpg&qlt=%d&p=0&pid=hp", imageURLBase, "UHD", 100)

	// Add height, width and crop
	if height > 0 {
		finalURL += fmt.Sprintf("&h=%d", height)
	}

	if width > 0 {
		finalURL += fmt.Sprintf("&w=%d", width)
	}

	if crop {
		finalURL += "&c=7" // smart crop (should crop keeping the most interesting part of the image)
		//finalURL += "&c=4" // standard crop (should crop from the center)
	}

	return finalURL, nil
}
