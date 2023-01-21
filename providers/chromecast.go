package providers

import (
	"fmt"
	"math/rand"

	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// Stuff to remove from the script element, that isn't part of the json
const removeBefore = "angular.module('home.constants', []). constant('fakeTimestamp',  null ). constant('initialStateJson', JSON.parse('"
const removeAfter = "')). constant('isAndroidTv',  false ). constant('isTextPromoEnabled',  false ). constant('isImaxVoiceQueryEnabled',  false ). constant('isManhattan',  false ). constant('isTouchNavigationEnabled',  false ). constant('isVoicePromoLanguage',  true ). constant('isOffersPromoEnabled',  false ). constant('imaxClientLogLevel', 'WARNING'). constant('isGFiberMessagingEnabled',  false ). constant('isImageBytesLoggingEnabled',  false ). constant('isGoogleSansUiEnabled',  false ). constant('isCacMessageThroughCcsEnabled',  true ). constant('isRemoteControlEnabled',  true ). constant('isPreloadingPhotosEnabled',  false ). constant('isSmartCropEnabled',  false ). constant('isDecodeAheadEnabled',  false ). constant('maxPreloadedTopics',  2.0 ). constant('isPauseOnHiddenEnabled',  false ). constant('sendAssistantMessengerContext',  false ). constant('fakeWeatherInfoJson', JSON.parse('null')). constant('minImageFailureDelayMs',  500.0 ). constant('maxImageFailureDelayMs',  10000.0 ). constant('imageFailureBackoffMultiplier',  2.0 );"

// Replaces escaped characters with their non escaped version
var unescaper = strings.NewReplacer("\\x5b", "[", "\\x22", "\"", "\\/", "/", "\\x5d", "]", "\\x27", "'")

// Gets a random photo link from "https://clients3.google.com/cast/chromecast/home/"
func getImageUrlChromecast(index int) (string, error) {
	// Parse the webpage
	doc, err := parseWebpageGoquery("https://clients3.google.com/cast/chromecast/home/")
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

func ChromecastWallpaper(height int, width int, crop int, portrait bool) (string, error) {
	// Gets a random url from the list of 50 urls in the chromecast homepage
	link, err := getImageUrlChromecast(rand.Intn(50))
	if err != nil {
		return "", err
	}
	var parameters string

	// add crop
	switch crop {
	// blind ratio
	case 1:
		parameters = fmt.Sprintf("w%d-h%d-c", width, height)
		break

	// smart ratio
	case 2:
		parameters = fmt.Sprintf("w%d-h%d-p", width, height)
		break

	default:
		if portrait {
			parameters = fmt.Sprintf("s%d", height)
		} else {
			parameters = fmt.Sprintf("s%d", width)
		}
		break
	}

	// Replace default parameters with modified ones
	finalLink := fmt.Sprintf("%s=%s", strings.Split(link, "=")[0], parameters)
	//fmt.Println("Image URL:", finalLink)

	return finalLink, nil
}
