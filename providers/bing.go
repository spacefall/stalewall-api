package providers

import (
	"fmt"
	"math/rand"

	"github.com/tidwall/gjson"
)

func BingWallpaper(market string, resolution string, quality int, height int, width int, crop int) (string, error) {
	// Generates the bing link, randomizing the index and market
	bingLink := fmt.Sprintf("https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=1&mkt=%s", rand.Intn(7), market)
	//fmt.Print("JSON URL:", bingLink)

	// Get Json
	content, err := parseWebpageIO(bingLink)
	if err != nil {
		return "", err
	}

	// Parse urlbase
	imageUrlBase := gjson.GetBytes(content, "images.0.urlbase").String()

	// create imageurl
	finalLink := fmt.Sprintf("http://bing.com%s_%s.jpg&qlt=%d&p=0&pid=hp", imageUrlBase, resolution, quality)

	// add height and width parameters if both are higher than 0
	if height > 0 {
		finalLink += fmt.Sprintf("&h=%d", height)
	}
	if width > 0 {
		finalLink += fmt.Sprintf("&w=%d", width)
	}

	// add crop
	switch crop {
	// blind ratio
	case 1:
		finalLink += "&c=4"
		break

	// smart ratio
	case 2:
		finalLink += "&c=7"
		break

	default:
		break
	}

	//fmt.Println("Final URL:", finalLink)

	return finalLink, nil
}
