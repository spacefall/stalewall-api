package main

// why am i trying to code in a language i don't even know

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/reujab/wallpaper"
	dnld "github.com/spaceox/gowall/downloaders"
	lib "github.com/spaceox/gowall/lib"
	"github.com/spf13/viper"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	lib.LoadConfig()
	// checks if pc is connected
	err := lib.Pong(viper.GetInt("ping.maxPingRetries"), viper.GetString("ping.pingRetrySleep"))
	check(err)

	// initializes rand with current time as seed
	rand.Seed(time.Now().Unix())
}

func main() {
	var pathToImage string
	var err error
	switch rand.Intn(2) {
	case 0:
		pathToImage, err = dnld.BingWallpaper(viper.GetStringSlice("bing.markets"), viper.GetString("bing.URLResolution"), viper.GetInt("bing.quality"), viper.GetInt("bing.height"), viper.GetInt("width"))
		check(err)
	case 1:
		pathToImage, err = dnld.ChromecastWallpaper(viper.GetString("chromecast.parameters"))
		check(err)
	}

	// applies wallpaper
	if viper.GetBool("wallpaper.applyafterdownload") {
		log.Println("applying image, this could take a while")
		err = wallpaper.SetFromFile(pathToImage)
		check(err)
		log.Println("image applied!")
	}

	// deletes wallpaper
	if viper.GetBool("wallpaper.deleteafterapply") {
		err = os.Remove(pathToImage)
		check(err)
		log.Println("image deleted")
	}

	log.Println("done")
}
