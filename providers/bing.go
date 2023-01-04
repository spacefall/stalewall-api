package providers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/tidwall/gjson"
)

func BingWallpaper(market string, resolution string, quality string, height string, width string) (string, error) {
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
	finalLink := fmt.Sprintf("http://bing.com%s_%s.jpg&qlt=%s", imageUrlBase, resolution, quality)

	// Convert height and width to int
	hint, err := strconv.Atoi(height)
	if err != nil {
		return "", err
	}
	wint, err := strconv.Atoi(width)
	if err != nil {
		return "", err
	}
	// add height and width parameters if both are higher than 0
	if hint > 0 {
		finalLink += fmt.Sprintf("&h=%s", height)
	}
	if wint > 0 {
		finalLink += fmt.Sprintf("&w=%s", width)
	}
	//fmt.Println("Final URL:", finalLink)

	return finalLink, nil
}
