package handler

import (
	"hash/maphash"
	"math/rand"
	"strconv"
	"strings"

	providers "github.com/spaceox/stalewall/providers"

	"fmt"
	"net/http"
)

var (
	bMarkets = [38]string{"es-AR", "en-AU", "de-AT", "nl-BE", "fr-BE", "pt-BR", "en-CA", "fr-CA", "es-CL", "da-DK", "fi-FI", "fr-FR", "de-DE", "zh-HK", "en-IN", "en-ID", "it-IT", "ja-JP", "ko-KR", "en-MY", "es-MX", "nl-NL", "en-NZ", "no-NO", "zh-CN", "pl-PL", "en-PH", "ru-RU", "en-ZA", "es-ES", "sv-SE", "fr-CH", "de-CH", "zh-TW", "tr-TR", "en-GB", "en-US", "es-US"}
	sLocales = [348]string{"af-NA", "am-ET", "ar-AE", "ar-BH", "ar-DJ", "ar-DZ", "ar-EG", "ar-EH", "ar-ER", "ar-IL", "ar-IQ", "ar-JO", "ar-KM", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-MR", "ar-OM", "ar-PS", "ar-QA", "ar-SA", "ar-SO", "ar-SS", "ar-TD", "ar-TN", "ar-YE", "as-IN", "az-AZ", "be-BY", "bg-BG", "bn-BD", "bn-IN", "bs-BA", "ca-AD", "ca-ES", "ca-FR", "ca-IT", "cs-CZ", "cy-GB", "da-DK", "da-GL", "de-AT", "de-BE", "de-CH", "de-DE", "de-IT", "de-LI", "de-LU", "el-CY", "el-GR", "en-AG", "en-AI", "en-AS", "en-AT", "en-AU", "en-BB", "en-BE", "en-BI", "en-BM", "en-BS", "en-BW", "en-BZ", "en-CA", "en-CC", "en-CH", "en-CK", "en-CM", "en-CX", "en-CY", "en-DE", "en-DG", "en-DM", "en-ER", "en-FJ", "en-FK", "en-FM", "en-GB", "en-GD", "en-GG", "en-GH", "en-GI", "en-GM", "en-GU", "en-GY", "en-HK", "en-ID", "en-IE", "en-IM", "en-IN", "en-IO", "en-JE", "en-JM", "en-KE", "en-KI", "en-KN", "en-KY", "en-LC", "en-LR", "en-LS", "en-MG", "en-MH", "en-MO", "en-MP", "en-MS", "en-MT", "en-MU", "en-MW", "en-MY", "en-NA", "en-NF", "en-NG", "en-NL", "en-NR", "en-NU", "en-NZ", "en-PG", "en-PH", "en-PK", "en-PN", "en-PR", "en-PW", "en-RW", "en-SB", "en-SC", "en-SE", "en-SG", "en-SH", "en-SI", "en-SL", "en-SS", "en-SX", "en-SZ", "en-TC", "en-TK", "en-TO", "en-TT", "en-TV", "en-TZ", "en-UG", "en-UM", "en-US", "en-VC", "en-VG", "en-VI", "en-VU", "en-WS", "en-ZA", "en-ZM", "en-ZW", "es-AR", "es-BO", "es-BR", "es-BZ", "es-CL", "es-CO", "es-CR", "es-DO", "es-EA", "es-EC", "es-ES", "es-GQ", "es-GT", "es-HN", "es-IC", "es-MX", "es-NI", "es-PA", "es-PE", "es-PH", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et-EE", "eu-ES", "fa-AF", "fi-FI", "fo-DR", "fo-FO", "fr-BE", "fr-BF", "fr-BI", "fr-BJ", "fr-BL", "fr-CA", "fr-CD", "fr-CF", "fr-CG", "fr-CH", "fr-CI", "fr-CM", "fr-DJ", "fr-DZ", "fr-FR", "fr-GA", "fr-GF", "fr-GN", "fr-GP", "fr-GQ", "fr-HT", "fr-KM", "fr-LU", "fr-MA", "fr-MC", "fr-MF", "fr-MG", "fr-ML", "fr-MQ", "fr-MR", "fr-MU", "fr-NC", "fr-NE", "fr-PF", "fr-PM", "fr-RE", "fr-RW", "fr-SC", "fr-SN", "fr-TD", "fr-TG", "fr-TN", "fr-VU", "fr-WF", "fr-YT", "gd-GB", "gl-ES", "gu-IN", "ha-GH", "ha-NE", "ha-NG", "he-IL", "hi-IN", "hr-BA", "hr-HR", "hu-HU", "hy-AM", "id-ID", "ig-NG", "is-IS", "it-CH", "it-IT", "it-SM", "it-VA", "ja-JP", "ka-GE", "kk-KZ", "kl-GL", "km-KH", "kn-IN", "ko-KR", "ku-IQ", "ky-KG", "lb-LU", "lo-LA", "lt-LT", "lv-LT", "lv-LV", "mk-MK", "ml-IN", "mn-CN", "mn-MN", "mr-IN", "ms-BN", "mt-MT", "nb-NO", "nb-SJ", "ne-IN", "ne-NP", "nl-AW", "nl-BE", "nl-BQ", "nl-CW", "nl-NL", "nl-SR", "nl-SX", "or-IN", "pa-IN", "pa-PK", "pl-PL", "ps-AF", "ps-PK", "pt-AO", "pt-BR", "pt-CH", "pt-CV", "pt-GQ", "pt-GW", "pt-LU", "pt-MO", "pt-MZ", "pt-PT", "pt-ST", "pt-TL", "ro-MD", "ro-RO", "ru-BY", "ru-KG", "ru-KZ", "ru-MD", "ru-RU", "rw-RW", "sd-IN", "sd-PK", "si-LK", "sk-SK", "sl-SI", "sl-SL", "sq-AL", "sq-MK", "sq-XK", "sr-BA", "sr-CS", "sr-ME", "sr-RS", "sr-XK", "sv-AX", "sv-FI", "sv-SE", "sw-CD", "sw-KE", "sw-TZ", "sw-UG", "ta-IN", "ta-LK", "te-IN", "th-TH", "ti-ET", "tk-TM", "tn-BW", "tr-CY", "tr-TR", "ug-CN", "ur-IN", "ur-PK", "uz-AF", "uz-UZ", "vi-VN", "wo-SN", "yo-BJ", "yo-NG", "zh-CN", "zh-HK", "zh-MO", "zh-SG", "zh-TW"}
)

type settings struct {
	bMkt     string
	bRes     string
	bQlt     int
	height   int
	width    int
	crop     int
	sLocale  string
	sPortrait bool
}

func init() {
	// Initializes rand with current time as seed
	rand.Seed(int64(new(maphash.Hash).Sum64()))
}

// Made this to make the code a bit cleaner
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func setup(r *http.Request) *settings {

	config := settings{
		bRes:     "UHD",
		bQlt:     90,
		height:   1080,
		width:    1920,
		crop:     0,
		sPortrait: false,
		sLocale:  sLocales[rand.Intn(len(sLocales))],
		bMkt:     bMarkets[rand.Intn(len(bMarkets))],
	}

	// gets an image that is as close as it is possible to a raw image
	if r.URL.Query().Has("raw") {
		// setting height and width to zero
		config.width, config.height = 0, 0
		// set bing quality to 100
		config.bQlt = 100
	}

	// if raw isn't specified but res is, the image will be requested in that resolution
	if r.URL.Query().Has("res") && !r.URL.Query().Has("raw") {
		// Gets res query from url
		resSplit := strings.Split(r.URL.Query().Get("res"), "x")

		// if resSplit[0] can be translated into an int and is bigger than zero, then set width to resSplit[0] as int
		if wint, err := strconv.Atoi(resSplit[0]); err == nil && wint > 0 {
			config.width = wint
		}
		// same thing with height
		if hint, err := strconv.Atoi(resSplit[1]); err == nil && hint > 0 {
			config.height = hint
		}

		// set if spotlight should use portrait images
		if config.height > config.width {
			config.sPortrait = true
		} else {
			config.sPortrait = false
		}
	}

	// crop
	// smart crop - finds the most interesting part of the image and crops to that
	// can be enabled only if raw isn't enabled
    if r.URL.Query().Has("scrop") && r.URL.Query().Has("res") && !r.URL.Query().Has("raw") {
		config.crop = 2
	}

	// standard crop - crops from the center
	// can be enabled only if raw and smart crop isn't enabled
    if r.URL.Query().Has("crop") && r.URL.Query().Has("res") && !r.URL.Query().Has("scrop") && !r.URL.Query().Has("raw") {
		config.crop = 1
	}

	// sets a market/locale for bing and spotlight
	if r.URL.Query().Has("mkt") {
		config.bMkt = r.URL.Query().Get("mkt")
		config.sLocale = r.URL.Query().Get("mkt")
	}

	// sets the filename resolution for bing
	if r.URL.Query().Has("bFilenameRes") {
		config.bRes = r.URL.Query().Get("bFilenameRes")
	}

	// sets bing quality parameter
	// can be enabled only if raw isn't enabled
	if r.URL.Query().Has("bQlt") && !r.URL.Query().Has("raw") {
		if qltint, err := strconv.Atoi(r.URL.Query().Get("bQlt")); err == nil {
			config.bQlt = qltint
		}
	}

	// forces spotlight image in portrait
	if r.URL.Query().Has("portrait") {
		config.sPortrait = true
	}

	return &config

}

func getWall(config *settings) string {
	// Initializing variables here to access the variables inside the switch, ouside
	switch rand.Intn(3) {
	case 0:
		// Bing wallpaper
		imageUrl, err := providers.BingWallpaper(config.bMkt, config.bRes, config.bQlt, config.height, config.width, config.crop)
		check(err)
		return createJson("bing", imageUrl)
	case 1:
		// Chromecast wallpaper
		imageUrl, err := providers.ChromecastWallpaper(config.height, config.width, config.crop)
		check(err)
		return createJson("chromecast", imageUrl)
	case 2:
		// Spotlight wallpaper
		imageUrl, err := providers.SpotlightWallpaper(config.sLocale, config.sPortrait)
		check(err)
		return createJson("spotlight", imageUrl)
	}
	panic("getWall: no value returned")

}

func createJson(source string, url string) string {
	return fmt.Sprintf(`{"source": %q, "url": %q}`, source, url)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	config := setup(r)
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := fmt.Fprintf(w, "%s", getWall(config))
	check(err)
}
