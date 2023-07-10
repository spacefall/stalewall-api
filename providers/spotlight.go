package providers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/tidwall/gjson"
)

// Gets a wallpaper from Windows Spotlight
func SpotlightWallpaper(locale string, portrait bool) (string, error) {
	// Generates the spotlight url
	// this is the old url used, i think it caused issues sometimes
	//apiURL := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=338387&fmt=json&ua=WindowsShellClient&cdm=1&pl=%s&ctry=%s", locale, strings.ToLower(strings.Split(locale, "-")[1]))
	apiURL := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=209567&fmt=json&rafb=0&cdm=1&lo=80217&pl=%s&lc=%s&ctry=%s", locale, locale, strings.Split(locale, "-")[1]) + "&ua=WindowsShellClient%252F0"

	// Get JSON
	content, err := parseWebpageIO(apiURL)
	if err != nil {
		return "", err
	}

	// Parse JSON
	itemJSON := gjson.GetBytes(content, fmt.Sprintf("batchrsp.items.%d.item", rand.Intn(3)))
	innerJSON := gjson.Parse(itemJSON.String())

	// Return urlbase
	if portrait {
		return innerJSON.Get("ad.image_fullscreen_001_portrait.u").String(), nil
	} else {
		return innerJSON.Get("ad.image_fullscreen_001_landscape.u").String(), nil
	}

}
