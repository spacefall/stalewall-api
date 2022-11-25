package lib

// everything that gets loaded from the web

import (
	//"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// DownloadFile will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
//
// Shamelessly copied from: https://gist.github.com/cnu/026744b1e86c6d9e22313d06cba4c2e9
func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Gets urlbase from provided Bing image archive url
func GetImageBaseURLBing(link string) (string, error) {
	// Loads the webpage
	res, err := http.Get(link)
	if err != nil {
		return "", err
	}
	// Gets the json
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	// Return urlbase
	return gjson.GetBytes(content, "images.0.urlbase").String(), nil
}

// Stuff to remove from the script element, that isn't part of the json
const removeBefore = "angular.module('home.constants', []). constant('fakeTimestamp',  null ). constant('initialStateJson', JSON.parse('"
const removeAfter = "')). constant('isAndroidTv',  false ). constant('isTextPromoEnabled',  false ). constant('isImaxVoiceQueryEnabled',  false ). constant('isManhattan',  false ). constant('isTouchNavigationEnabled',  false ). constant('isVoicePromoLanguage',  true ). constant('isOffersPromoEnabled',  false ). constant('imaxClientLogLevel', 'WARNING'). constant('isGFiberMessagingEnabled',  false ). constant('isImageBytesLoggingEnabled',  false ). constant('isGoogleSansUiEnabled',  false ). constant('isCacMessageThroughCcsEnabled',  true ). constant('isRemoteControlEnabled',  true ). constant('isPreloadingPhotosEnabled',  false ). constant('isSmartCropEnabled',  false ). constant('isDecodeAheadEnabled',  false ). constant('maxPreloadedTopics',  2.0 ). constant('isPauseOnHiddenEnabled',  false ). constant('sendAssistantMessengerContext',  false ). constant('fakeWeatherInfoJson', JSON.parse('null')). constant('minImageFailureDelayMs',  500.0 ). constant('maxImageFailureDelayMs',  10000.0 ). constant('imageFailureBackoffMultiplier',  2.0 );"

// Gets a random photo link from "https://clients3.google.com/cast/chromecast/home/"
func GetImageUrlChromecast(index int) (string, error) {
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

// Gets html name from link
func GetChromecastImageTitle(url string) (string, error) {
	// Loads the webpage
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	headers := res.Header.Values("content-disposition")[0]
	return strings.Trim(strings.Split(headers, "=")[1], "\""), nil
}
