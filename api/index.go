package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	providers "github.com/spaceox/stalewall/providers"
)

var (
	bMarkets = [27]string{"es-AR", "en-AU", "de-AT", "nl-BE", "fr-BE", "pt-BR", "en-CA", "fr-CA", "da-DK", "fi-FI", "fr-FR", "de-DE", "zh-HK", "en-IN", "en-ID", "it-IT", "ja-JP", "ko-KR", "zh-CN", "pl-PL", "ru-RU", "es-ES", "sv-SE", "tr-TR", "en-GB", "en-US", "es-US"}
	sLocales = [348]string{"af-NA", "am-ET", "ar-AE", "ar-BH", "ar-DJ", "ar-DZ", "ar-EG", "ar-EH", "ar-ER", "ar-IL", "ar-IQ", "ar-JO", "ar-KM", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-MR", "ar-OM", "ar-PS", "ar-QA", "ar-SA", "ar-SO", "ar-SS", "ar-TD", "ar-TN", "ar-YE", "as-IN", "az-AZ", "be-BY", "bg-BG", "bn-BD", "bn-IN", "bs-BA", "ca-AD", "ca-ES", "ca-FR", "ca-IT", "cs-CZ", "cy-GB", "da-DK", "da-GL", "de-AT", "de-BE", "de-CH", "de-DE", "de-IT", "de-LI", "de-LU", "el-CY", "el-GR", "en-AG", "en-AI", "en-AS", "en-AT", "en-AU", "en-BB", "en-BE", "en-BI", "en-BM", "en-BS", "en-BW", "en-BZ", "en-CA", "en-CC", "en-CH", "en-CK", "en-CM", "en-CX", "en-CY", "en-DE", "en-DG", "en-DM", "en-ER", "en-FJ", "en-FK", "en-FM", "en-GB", "en-GD", "en-GG", "en-GH", "en-GI", "en-GM", "en-GU", "en-GY", "en-HK", "en-ID", "en-IE", "en-IM", "en-IN", "en-IO", "en-JE", "en-JM", "en-KE", "en-KI", "en-KN", "en-KY", "en-LC", "en-LR", "en-LS", "en-MG", "en-MH", "en-MO", "en-MP", "en-MS", "en-MT", "en-MU", "en-MW", "en-MY", "en-NA", "en-NF", "en-NG", "en-NL", "en-NR", "en-NU", "en-NZ", "en-PG", "en-PH", "en-PK", "en-PN", "en-PR", "en-PW", "en-RW", "en-SB", "en-SC", "en-SE", "en-SG", "en-SH", "en-SI", "en-SL", "en-SS", "en-SX", "en-SZ", "en-TC", "en-TK", "en-TO", "en-TT", "en-TV", "en-TZ", "en-UG", "en-UM", "en-US", "en-VC", "en-VG", "en-VI", "en-VU", "en-WS", "en-ZA", "en-ZM", "en-ZW", "es-AR", "es-BO", "es-BR", "es-BZ", "es-CL", "es-CO", "es-CR", "es-DO", "es-EA", "es-EC", "es-ES", "es-GQ", "es-GT", "es-HN", "es-IC", "es-MX", "es-NI", "es-PA", "es-PE", "es-PH", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et-EE", "eu-ES", "fa-AF", "fi-FI", "fo-DR", "fo-FO", "fr-BE", "fr-BF", "fr-BI", "fr-BJ", "fr-BL", "fr-CA", "fr-CD", "fr-CF", "fr-CG", "fr-CH", "fr-CI", "fr-CM", "fr-DJ", "fr-DZ", "fr-FR", "fr-GA", "fr-GF", "fr-GN", "fr-GP", "fr-GQ", "fr-HT", "fr-KM", "fr-LU", "fr-MA", "fr-MC", "fr-MF", "fr-MG", "fr-ML", "fr-MQ", "fr-MR", "fr-MU", "fr-NC", "fr-NE", "fr-PF", "fr-PM", "fr-RE", "fr-RW", "fr-SC", "fr-SN", "fr-TD", "fr-TG", "fr-TN", "fr-VU", "fr-WF", "fr-YT", "gd-GB", "gl-ES", "gu-IN", "ha-GH", "ha-NE", "ha-NG", "he-IL", "hi-IN", "hr-BA", "hr-HR", "hu-HU", "hy-AM", "id-ID", "ig-NG", "is-IS", "it-CH", "it-IT", "it-SM", "it-VA", "ja-JP", "ka-GE", "kk-KZ", "kl-GL", "km-KH", "kn-IN", "ko-KR", "ku-IQ", "ky-KG", "lb-LU", "lo-LA", "lt-LT", "lv-LT", "lv-LV", "mk-MK", "ml-IN", "mn-CN", "mn-MN", "mr-IN", "ms-BN", "mt-MT", "nb-NO", "nb-SJ", "ne-IN", "ne-NP", "nl-AW", "nl-BE", "nl-BQ", "nl-CW", "nl-NL", "nl-SR", "nl-SX", "or-IN", "pa-IN", "pa-PK", "pl-PL", "ps-AF", "ps-PK", "pt-AO", "pt-BR", "pt-CH", "pt-CV", "pt-GQ", "pt-GW", "pt-LU", "pt-MO", "pt-MZ", "pt-PT", "pt-ST", "pt-TL", "ro-MD", "ro-RO", "ru-BY", "ru-KG", "ru-KZ", "ru-MD", "ru-RU", "rw-RW", "sd-IN", "sd-PK", "si-LK", "sk-SK", "sl-SI", "sl-SL", "sq-AL", "sq-MK", "sq-XK", "sr-BA", "sr-CS", "sr-ME", "sr-RS", "sr-XK", "sv-AX", "sv-FI", "sv-SE", "sw-CD", "sw-KE", "sw-TZ", "sw-UG", "ta-IN", "ta-LK", "te-IN", "th-TH", "ti-ET", "tk-TM", "tn-BW", "tr-CY", "tr-TR", "ug-CN", "ur-IN", "ur-PK", "uz-AF", "uz-UZ", "vi-VN", "wo-SN", "yo-BJ", "yo-NG", "zh-CN", "zh-HK", "zh-MO", "zh-SG", "zh-TW"}
)

type settings struct {
	bMkt      string
	bRes      string
	bQlt      int
	height    int
	width     int
	crop      int
	sLocale   string
	sPortrait bool
}

func parseQueries(r *http.Request) (*settings, error) {

	config := settings{
		bRes:      "UHD",
		bQlt:      100,
		bMkt:      bMarkets[rand.Intn(len(bMarkets))],
		sPortrait: false,
		sLocale:   sLocales[rand.Intn(len(sLocales))],
		height:    0,
		width:     0,
		crop:      0,
	}

	// if res is specified, the image will be requested in that resolution
	if r.URL.Query().Has("res") {

		// Gets res query from url
		resSplit := strings.Split(r.URL.Query().Get("res"), "x")

		// if resSplit[0] can be translated into an int and is bigger than zero, then set width to resSplit[0] as int
		wint, err := strconv.Atoi(resSplit[0])
		if err == nil && wint > 0 {
			config.width = wint
		} else if err != nil {
			return nil, err
		}

		// same thing with height
		hint, err := strconv.Atoi(resSplit[1])
		if err == nil && hint > 0 {
			config.height = hint
		} else if err != nil {
			return nil, err
		}

		// set if spotlight should use portrait images
		if config.height > config.width {
			config.sPortrait = true
		}
	}

	// crop
	// this just asks the source to crop the image, it doesn't crop the image
	if r.URL.Query().Has("res") {
		if r.URL.Query().Has("scrop") { // smart crop - finds the most interesting part of the image and crops to that
			config.crop = 2
		} else if r.URL.Query().Has("crop") { // standard crop - crops from the center
			config.crop = 1
		}
	}

	// sets a market/locale for bing and spotlight
	if r.URL.Query().Has("mkt") {
		config.bMkt = r.URL.Query().Get("mkt")
		config.sLocale = r.URL.Query().Get("mkt")
	}

	// sets the resolution for bing by changing the filename resolution
	if r.URL.Query().Has("bRes") {
		config.bRes = r.URL.Query().Get("bRes")
	}

	// sets bing quality parameter
	if r.URL.Query().Has("bQlt") {
		qltint, err := strconv.Atoi(r.URL.Query().Get("bQlt"))
		if err == nil {
			config.bQlt = qltint
		} else {
			return nil, err
		}
	}

	// forces spotlight image in portrait
	if r.URL.Query().Has("portrait") {
		config.sPortrait = true
	}

	return &config, nil
}

func getWall(config *settings) ([]byte, error) {
	// Initializing variables here to access the variables inside the switch, ouside
	var (
		imageUrl string
		provider string
		err      error
	)
	switch rand.Intn(3) {
	case 0:
		// Bing wallpaper
		provider = "bing"
		imageUrl, err = providers.BingWallpaper(config.bMkt, config.bRes, config.bQlt, config.height, config.width, config.crop)
	case 1:
		// Chromecast wallpaper
		provider = "chromecast"
		imageUrl, err = providers.ChromecastWallpaper(config.height, config.width, config.crop)
	case 2:
		// Spotlight wallpaper
		provider = "spotlight"
		imageUrl, err = providers.SpotlightWallpaper(config.sLocale, config.sPortrait)
	default:
		return nil, errors.New("index out of range in switch on function getWall")
	}

	if err != nil {
		return nil, err
	}

	return createJSON(provider, imageUrl)
}

// creates a json, it's put here to make json modifications more streamlined
// TODO: break the client by changing source to provider
func createJSON(provider, url string) ([]byte, error) {
	data := struct {
		Provider string `json:"source"`
		URL      string `json:"url"`
	}{
		Provider: provider,
		URL:      url,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "You can't like do that")
		return
	}

	config, err := parseQueries(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wall, err := getWall(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// identifies the page as a json
	w.Header().Set("Content-Type", "application/json")
	// allows cors
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// write the json
	w.Write(wall)

	// writes the json to the page
	/* _, err = fmt.Fprint(w, wall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} */
}
