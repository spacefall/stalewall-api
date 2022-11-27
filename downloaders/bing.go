package downloaders

import (
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	lib "github.com/spaceox/gowall/lib"
	"github.com/tidwall/gjson"
)

// Gets urlbase from provided Bing image archive url
func getImageBaseURLBing(link string) (string, error) {
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

	// Return urlbase
	return gjson.GetBytes(content, "images.0.urlbase").String(), nil
}

func BingWallpaper(markets []string, URLResolution string, quality int, height int, width int) (string, error) {
	// Generates the bing link, randomizing the index and market
	bingLink := "https://www.bing.com/HPImageArchive.aspx?format=js&idx=" + strconv.Itoa(rand.Intn(7)) + "&n=1&mkt=" + markets[rand.Intn(len(markets))]
	lib.LogInColor.Info("JSON URL:", bingLink)

	// Get image baseurl
	imageUrlBase, err := getImageBaseURLBing(bingLink)
	if err != nil {
		return "", err
	}

	// Define filename
	imageName := strings.Split(imageUrlBase, "OHR.")[1] + ".jpg"

	// create imageurl
	finalLink := "http://bing.com" + imageUrlBase + "_" + URLResolution + ".jpg&qlt=" + strconv.Itoa(quality)
	if height > 0 && width > 0 {
		finalLink += "&h=" + strconv.Itoa(height) + "&w=" + strconv.Itoa(width)
	}
	lib.LogInColor.Info("Final URL:", finalLink)

	// download
	err = lib.DownloadFile(finalLink, imageName)
	if err != nil {
		return "", err
	}
	// Gets image path
	return lib.GetFinalPath(imageName)
}
