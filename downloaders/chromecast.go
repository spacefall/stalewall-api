package downloaders

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	lib "github.com/spaceox/gowall/lib"
)

func ChromecastWallpaper(parameters string) (string, error) {
	link, err := lib.GetImageUrlChromecast(rand.Intn(49))
	if err != nil {
		return "", err
	}
	log.Println("imageBaseURL:", link)

	finalLink := strings.Split(link, "=")[0] + "=" + parameters
	log.Println("url:", finalLink)

	imageName, err := lib.GetChromecastImageTitle(finalLink)
	if err != nil {
		return "", err
	}
	log.Println("imageName:", imageName)

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
