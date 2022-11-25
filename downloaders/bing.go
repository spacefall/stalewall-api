package downloaders

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	lib "github.com/spaceox/gowall/lib"
)

func BingWallpaper(markets []string, URLResolution string, quality int, height int, width int) (string, error) {
	// generates the bing link, randomizing the index and market
	bingLink := "https://www.bing.com/HPImageArchive.aspx?format=js&idx=" + strconv.Itoa(rand.Intn(7)) + "&n=1&mkt=" + markets[rand.Intn(len(markets))]
	log.Println("url:", bingLink)

	// get image baseurl
	imageUrlBase, err := lib.GetImageBaseURLBing(bingLink)
	if err != nil {
		return "", err
	}
	log.Println("imageUrlBase:", imageUrlBase)

	// create imagename
	imageName := strings.Split(imageUrlBase, "OHR.")[1] + ".jpg"
	log.Println("imageName:", imageName)

	// create imageurl
	finalLink := "http://bing.com" + imageUrlBase + "_" + URLResolution + ".jpg&qlt=" + strconv.Itoa(quality) + "&h=" + strconv.Itoa(height) + "&w=" + strconv.Itoa(width)
	log.Println("imageurl:", finalLink)

	// download
	err = lib.DownloadFile(finalLink, imageName)
	if err != nil {
		return "", err
	}
	// gets path to wallpaper
	workdir, err := os.Getwd()
	pathToImage := filepath.Join(workdir, imageName)
	if err != nil {
		return "", err
	}
	log.Println("path:", pathToImage)
	return pathToImage, nil
}
