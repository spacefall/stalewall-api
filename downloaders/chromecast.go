package downloaders

import (
	"io"
	"math/rand"
	"net/http"
	"os"

	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"

	lib "github.com/spaceox/gowall/lib"
)

// Gets a random photo link from "https://clients3.google.com/cast/chromecast/home/"
func getImageUrlChromecast(index int) (string, error) {
	// Stuff to remove from the script element, that isn't part of the json
	const removeBefore = "angular.module('home.constants', []). constant('fakeTimestamp',  null ). constant('initialStateJson', JSON.parse('"
	const removeAfter = "')). constant('isAndroidTv',  false ). constant('isTextPromoEnabled',  false ). constant('isImaxVoiceQueryEnabled',  false ). constant('isManhattan',  false ). constant('isTouchNavigationEnabled',  false ). constant('isVoicePromoLanguage',  true ). constant('isOffersPromoEnabled',  false ). constant('imaxClientLogLevel', 'WARNING'). constant('isGFiberMessagingEnabled',  false ). constant('isImageBytesLoggingEnabled',  false ). constant('isGoogleSansUiEnabled',  false ). constant('isCacMessageThroughCcsEnabled',  true ). constant('isRemoteControlEnabled',  true ). constant('isPreloadingPhotosEnabled',  false ). constant('isSmartCropEnabled',  false ). constant('isDecodeAheadEnabled',  false ). constant('maxPreloadedTopics',  2.0 ). constant('isPauseOnHiddenEnabled',  false ). constant('sendAssistantMessengerContext',  false ). constant('fakeWeatherInfoJson', JSON.parse('null')). constant('minImageFailureDelayMs',  500.0 ). constant('maxImageFailureDelayMs',  10000.0 ). constant('imageFailureBackoffMultiplier',  2.0 );"

	// Replaces escaped characters with their non escaped version
	var unescaper = strings.NewReplacer("\\x5b", "[", "\\x22", "\"", "\\/", "/", "\\x5d", "]", "\\x27", "'")

	// Loads the webpage
	res, err := http.Get("https://clients3.google.com/cast/chromecast/home/")
	if err != nil {
		return "", err
	}

	// Parse the webpage
	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	// Get script tag and load into 'hiddenJSON' the json
	hiddenJSON := ""
	doc.Find("script").Each(func(i int, element *goquery.Selection) {
		if element.Text() != "" {
			hiddenJSON = strings.Replace(strings.Replace(element.Text(), removeBefore, "", 1), removeAfter, "", 1)
		}
	})
	parsedJSON := gjson.Parse(unescaper.Replace(hiddenJSON))
	return strings.ReplaceAll(parsedJSON.Get("0").Get(strconv.Itoa(index)).Get("0").String(), "\\u003d", "="), nil
}

// Same function as the one in common.go, but adapted to extract filename from headers
func downloadChromecastImage(url string) (string, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Get filename
	contDispHeader := resp.Header.Values("content-disposition")[0]
	filename := strings.Trim(strings.Split(contDispHeader, "=")[1], "\"")

	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func ChromecastWallpaper(parameters string) (string, error) {
	// Gets a random url from the list of 50 urls in the chromecast homepage
	link, err := getImageUrlChromecast(rand.Intn(50))
	if err != nil {
		return "", err
	}

	// Replace default parameters with modified ones
	finalLink := strings.Split(link, "=")[0] + "=" + parameters
	lib.LogInColor.Info("Image URL:", finalLink)

	// Get filename and download image
	filename, err := downloadChromecastImage(finalLink)
	if err != nil {
		return "", err
	}

	// Gets path to image
	return lib.GetFinalPath(filename)
}
