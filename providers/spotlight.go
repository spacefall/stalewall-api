package providers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// SpotlightWallpaper - Gets a wallpaper from Windows Spotlight
func SpotlightWallpaper(locale string, portrait bool) (string, error) {
	// Generates the spotlight url
	// this is the old url used, i think it caused issues sometimes
	//apiURL := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=338387&fmt=json&ua=WindowsShellClient&cdm=1&pl=%s&ctry=%s", locale, strings.ToLower(strings.Split(locale, "-")[1]))

	apiURL := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=209567&fmt=json&rafb=0&cdm=1&lo=80217&pl=%s&lc=%s&ctry=%s", locale, locale, strings.Split(locale, "-")[1]) + "&ua=WindowsShellClient%252F0"

	// Get JSON
	res, err := http.Get(apiURL)
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

	// Parse JSON
	itemJSON := gjson.GetBytes(content, "batchrsp.items.0.item")
	innerJSON := gjson.Parse(itemJSON.String())

	// Return urlbase
	if portrait {
		return innerJSON.Get("ad.image_fullscreen_001_portrait.u").String(), nil
	} else {
		return innerJSON.Get("ad.image_fullscreen_001_landscape.u").String(), nil
	}

}
