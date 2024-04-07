package providers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

var (
	// Replaces escaped characters with their non escaped version
	unescaper = strings.NewReplacer("\\x5b", "[", "\\x22", "\"", "\\/", "/", "\\x5d", "]", "\\x27", "'")
)

func ChromecastWallpaper(height, width int, crop bool) (string, error) {
	// Chooses a random number 0-49 which will be the photo selected from the list
	imageIndex := rand.Intn(50)

	// Gets chromecast image page thing
	res, err := http.Get("https://clients3.google.com/cast/chromecast/home/")
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Checking that the status code is 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load response into a goquery document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	// Finds the script tag containing the json and strips the angular code from it
	hiddenJSON := doc.Find("body > script").Text()
	_, hiddenJSON, _ = strings.Cut(hiddenJSON, "JSON.parse('")
	hiddenJSON, _, _ = strings.Cut(hiddenJSON, "')).")

	// Unescaping the json, couldn't find a good way to unescape it without replacer
	parsedJSON := gjson.Parse(unescaper.Replace(hiddenJSON))

	// Grabs the image link
	imageURL := parsedJSON.Get(fmt.Sprintf("0.%d.0", imageIndex)).String()

	// Stitches the needed parameters
	finalURL, _, _ := strings.Cut(imageURL, "\\u003d")
	finalURL += fmt.Sprintf("=w%d-h%d", width, height)

	// Adds crop
	if crop {
		finalURL += "-p" // smart crop (should crop keeping the most interesting part of the image)
		// finalURL += "-c" // standard crop (should crop from the center of the image)
	}
	return finalURL, nil
}
