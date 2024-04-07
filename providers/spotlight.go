package providers

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

var sLocales = [348]string{"af-NA", "am-ET", "ar-AE", "ar-BH", "ar-DJ", "ar-DZ", "ar-EG", "ar-EH", "ar-ER", "ar-IL", "ar-IQ", "ar-JO", "ar-KM", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-MR", "ar-OM", "ar-PS", "ar-QA", "ar-SA", "ar-SO", "ar-SS", "ar-TD", "ar-TN", "ar-YE", "as-IN", "az-AZ", "be-BY", "bg-BG", "bn-BD", "bn-IN", "bs-BA", "ca-AD", "ca-ES", "ca-FR", "ca-IT", "cs-CZ", "cy-GB", "da-DK", "da-GL", "de-AT", "de-BE", "de-CH", "de-DE", "de-IT", "de-LI", "de-LU", "el-CY", "el-GR", "en-AG", "en-AI", "en-AS", "en-AT", "en-AU", "en-BB", "en-BE", "en-BI", "en-BM", "en-BS", "en-BW", "en-BZ", "en-CA", "en-CC", "en-CH", "en-CK", "en-CM", "en-CX", "en-CY", "en-DE", "en-DG", "en-DM", "en-ER", "en-FJ", "en-FK", "en-FM", "en-GB", "en-GD", "en-GG", "en-GH", "en-GI", "en-GM", "en-GU", "en-GY", "en-HK", "en-ID", "en-IE", "en-IM", "en-IN", "en-IO", "en-JE", "en-JM", "en-KE", "en-KI", "en-KN", "en-KY", "en-LC", "en-LR", "en-LS", "en-MG", "en-MH", "en-MO", "en-MP", "en-MS", "en-MT", "en-MU", "en-MW", "en-MY", "en-NA", "en-NF", "en-NG", "en-NL", "en-NR", "en-NU", "en-NZ", "en-PG", "en-PH", "en-PK", "en-PN", "en-PR", "en-PW", "en-RW", "en-SB", "en-SC", "en-SE", "en-SG", "en-SH", "en-SI", "en-SL", "en-SS", "en-SX", "en-SZ", "en-TC", "en-TK", "en-TO", "en-TT", "en-TV", "en-TZ", "en-UG", "en-UM", "en-US", "en-VC", "en-VG", "en-VI", "en-VU", "en-WS", "en-ZA", "en-ZM", "en-ZW", "es-AR", "es-BO", "es-BR", "es-BZ", "es-CL", "es-CO", "es-CR", "es-DO", "es-EA", "es-EC", "es-ES", "es-GQ", "es-GT", "es-HN", "es-IC", "es-MX", "es-NI", "es-PA", "es-PE", "es-PH", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et-EE", "eu-ES", "fa-AF", "fi-FI", "fo-DR", "fo-FO", "fr-BE", "fr-BF", "fr-BI", "fr-BJ", "fr-BL", "fr-CA", "fr-CD", "fr-CF", "fr-CG", "fr-CH", "fr-CI", "fr-CM", "fr-DJ", "fr-DZ", "fr-FR", "fr-GA", "fr-GF", "fr-GN", "fr-GP", "fr-GQ", "fr-HT", "fr-KM", "fr-LU", "fr-MA", "fr-MC", "fr-MF", "fr-MG", "fr-ML", "fr-MQ", "fr-MR", "fr-MU", "fr-NC", "fr-NE", "fr-PF", "fr-PM", "fr-RE", "fr-RW", "fr-SC", "fr-SN", "fr-TD", "fr-TG", "fr-TN", "fr-VU", "fr-WF", "fr-YT", "gd-GB", "gl-ES", "gu-IN", "ha-GH", "ha-NE", "ha-NG", "he-IL", "hi-IN", "hr-BA", "hr-HR", "hu-HU", "hy-AM", "id-ID", "ig-NG", "is-IS", "it-CH", "it-IT", "it-SM", "it-VA", "ja-JP", "ka-GE", "kk-KZ", "kl-GL", "km-KH", "kn-IN", "ko-KR", "ku-IQ", "ky-KG", "lb-LU", "lo-LA", "lt-LT", "lv-LT", "lv-LV", "mk-MK", "ml-IN", "mn-CN", "mn-MN", "mr-IN", "ms-BN", "mt-MT", "nb-NO", "nb-SJ", "ne-IN", "ne-NP", "nl-AW", "nl-BE", "nl-BQ", "nl-CW", "nl-NL", "nl-SR", "nl-SX", "or-IN", "pa-IN", "pa-PK", "pl-PL", "ps-AF", "ps-PK", "pt-AO", "pt-BR", "pt-CH", "pt-CV", "pt-GQ", "pt-GW", "pt-LU", "pt-MO", "pt-MZ", "pt-PT", "pt-ST", "pt-TL", "ro-MD", "ro-RO", "ru-BY", "ru-KG", "ru-KZ", "ru-MD", "ru-RU", "rw-RW", "sd-IN", "sd-PK", "si-LK", "sk-SK", "sl-SI", "sl-SL", "sq-AL", "sq-MK", "sq-XK", "sr-BA", "sr-CS", "sr-ME", "sr-RS", "sr-XK", "sv-AX", "sv-FI", "sv-SE", "sw-CD", "sw-KE", "sw-TZ", "sw-UG", "ta-IN", "ta-LK", "te-IN", "th-TH", "ti-ET", "tk-TM", "tn-BW", "tr-CY", "tr-TR", "ug-CN", "ur-IN", "ur-PK", "uz-AF", "uz-UZ", "vi-VN", "wo-SN", "yo-BJ", "yo-NG", "zh-CN", "zh-HK", "zh-MO", "zh-SG", "zh-TW"}

func SpotlightWallpaper(height, width int) (string, error) {
	locale := sLocales[rand.Intn(len(sLocales))]
	URL := fmt.Sprintf("https://arc.msn.com/v3/Delivery/Placement?pid=209567&fmt=json&rafb=0&cdm=1&lo=80217&pl=%s&lc=%s&ctry=%s", locale, locale, strings.Split(locale, "-")[1]) + "&ua=WindowsShellClient%252F0"

	// Get JSON
	res, err := http.Get(URL)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// Checking that the status code is 200
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Gets the response in bytes to decode it with gjson
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Parse the json from the response, then parse the json contained in batchrsp.items.0.item because it's encoded into a string for some reason
	itemJSON := gjson.GetBytes(content, "batchrsp.items.0.item")
	innerItemJSON := gjson.Parse(itemJSON.String())

	if height > width {
		return innerItemJSON.Get("ad.image_fullscreen_001_portrait.u").String(), nil
	} else {
		return innerItemJSON.Get("ad.image_fullscreen_001_landscape.u").String(), nil
	}

}
