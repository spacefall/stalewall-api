package providers

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// Gets urlbase from provided Bing image archive url
func getSpotlightImageLink(link string, portrait bool) (string, error) {
	// Get JSON
	content, err := parseWebpageIO(link)
	if err != nil {
		return "", err
	}

	// Parse JSON
	parsedJSON := gjson.Parse(gjson.GetBytes(content, "batchrsp.items.0.item").String())

	// Return urlbase
	if portrait {
		return parsedJSON.Get("ad.image_fullscreen_001_portrait.u").String(), nil
	} else {
		return parsedJSON.Get("ad.image_fullscreen_001_landscape.u").String(), nil
	}
}

func SpotlightWallpaper(locale string, portrait bool) (string, error) {
	// Generates the spotlight link
	apiLink := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=338387&fmt=json&ua=WindowsShellClient&cdm=1&pl=%s&ctry=%s", locale, strings.ToLower(strings.Split(locale, "-")[1]))
	//fmt.Println("JSON URL:", apiLink)

	// Get image baseurl
	imageUrlBase, err := getSpotlightImageLink(apiLink, portrait)
	if err != nil {
		return "", err
	}
	//fmt.Println("Final URL:", imageUrlBase)

	return imageUrlBase, nil
}
