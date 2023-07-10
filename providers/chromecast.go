package providers

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

// Stuff to remove from the script element, that isn't part of the json
const removeBefore = "angular.module('home.constants', []). constant('fakeTimestamp',  null ). constant('initialStateJson', JSON.parse('"
const removeAfter = "')). constant('isAndroidTv',  false ). constant('isTextPromoEnabled',  false ). constant('isImaxVoiceQueryEnabled',  false ). constant('isManhattan',  false ). constant('isTouchNavigationEnabled',  false ). constant('isVoicePromoLanguage',  true ). constant('isOffersPromoEnabled',  false ). constant('imaxClientLogLevel', 'WARNING'). constant('isGFiberMessagingEnabled',  false ). constant('isImageBytesLoggingEnabled',  false ). constant('isGoogleSansUiEnabled',  false ). constant('isCacMessageThroughCcsEnabled',  true ). constant('isRemoteControlEnabled',  true ). constant('isPreloadingPhotosEnabled',  false ). constant('isSmartCropEnabled',  false ). constant('isDecodeAheadEnabled',  false ). constant('maxPreloadedTopics',  2.0 ). constant('isPauseOnHiddenEnabled',  false ). constant('sendAssistantMessengerContext',  false ). constant('fakeWeatherInfoJson', JSON.parse('null')). constant('minImageFailureDelayMs',  500.0 ). constant('maxImageFailureDelayMs',  10000.0 ). constant('imageFailureBackoffMultiplier',  2.0 );"

var (
	// Removes the removeBefore and removeAfter strings from the script, leaving only the json
	debloater = strings.NewReplacer(removeAfter, "", removeBefore, "")
	// Replaces escaped characters with their non escaped version
	unescaper = strings.NewReplacer("\\x5b", "[", "\\x22", "\"", "\\/", "/", "\\x5d", "]", "\\x27", "'")
)

// Gets a random photo from "https://clients3.google.com/cast/chromecast/home/"
func ChromecastWallpaper(height, width, crop int) (string, error) {
	// initializing vars for later
	var (
		hiddenJSON string
		parameters string
	)

	// Chooses a random number 0-49 which will be the photo selected from the list
	imageIndex := rand.Intn(50)

	// Parse the webpage
	doc, err := parseWebpageGoquery("https://clients3.google.com/cast/chromecast/home/")
	if err != nil {
		return "", err
	}

	// Finds the script tag (that doesn't just url to a js file) and strips it of the unneeded angular code
	doc.Find("script").Each(func(i int, element *goquery.Selection) {
		if element.Text() != "" {
			hiddenJSON = debloater.Replace(element.Text())
		}
	})

	// The JSON gets unescaped and parsed
	// (couldn't find a way to unescape it without replacers)
	parsedJSON := gjson.Parse(unescaper.Replace(hiddenJSON))

	// The url is basically ready, just need to unescape the =
	imageURL := parsedJSON.Get(fmt.Sprintf("0.%d.0", imageIndex)).String()
	finalURL := strings.ReplaceAll(imageURL, "\\u003d", "=")

	// Adds crop
	switch crop {
	case 1:
		parameters = fmt.Sprintf("w%d-h%d-c", width, height) // blind ratio

	case 2:
		parameters = fmt.Sprintf("w%d-h%d-p", width, height) // smart ratio

	default:
		parameters = fmt.Sprintf("s%d", width)
	}

	// Replace default parameters with modified ones
	return fmt.Sprintf("%s=%s", strings.Split(finalURL, "=")[0], parameters), nil
}
