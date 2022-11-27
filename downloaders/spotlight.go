package downloaders

import (
	"io"
	"math/rand"
	"net/http"
	"strings"

	lib "github.com/spaceox/gowall/lib"
	"github.com/tidwall/gjson"
)

// Gets urlbase from provided Bing image archive url
func getSpotlightImageLink(link string, portrait bool) (string, error) {
	// Loads the webpage
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}

	// Gets the json
	content, err := io.ReadAll(res.Body)
	defer res.Body.Close()
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

func SpotlightWallpaper(locales []string, portrait bool) (string, error) {
	// Generates the spotlight link, randomizing the locale (and country)
	localeIndex := rand.Intn(len(locales))
	apiLink := "https://arc.msn.com/v3/Delivery/Placement?pid=338387&fmt=json&ua=WindowsShellClient&cdm=1&pl=" + locales[localeIndex] + "&ctry=" + strings.ToLower(strings.Split(locales[localeIndex], "-")[1])
	lib.LogInColor.Info("JSON URL:", apiLink)

	// Get image baseurl
	imageUrlBase, err := getSpotlightImageLink(apiLink, portrait)
	if err != nil {
		return "", err
	}
	lib.LogInColor.Info("Final URL:", imageUrlBase)

	// Define filename
	imageName := strings.Split(imageUrlBase, "=")[1] + ".jpg"

	// download
	err = lib.DownloadFile(imageUrlBase, imageName)
	if err != nil {
		return "", err
	}

	// Gets image path
	return lib.GetFinalPath(imageName)
}
