package main

import (
	"hash/maphash"
	"math/rand"
	"os"

	"github.com/reujab/wallpaper"
	dnld "github.com/spaceox/gowall/downloaders"
	lib "github.com/spaceox/gowall/lib"
	"github.com/spf13/viper"
)

// Made this to make the code a bit cleaner
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// Loads config from config.toml
	lib.LoadConfig()

	// Checks if pc is connected
	err := lib.Pong(
		viper.GetInt("ping.maxPingRetries"),
		viper.GetString("ping.pingRetrySleep"),
		viper.GetString("ping.timeout"),
	)
	check(err)

	// Initializes rand with current time as seed
	// TODO: Remove? when go 1.20 releases
	//rand.Seed(time.Now().Unix())
	rand.Seed(int64(new(maphash.Hash).Sum64()))
}

func main() {
	// Initializing variables here to access the variables inside the switch, ouside
	var (
		pathToImage string
		err         error
	)
	switch rand.Intn(2) {
	case 0:
		// Bing wallpaper
		pathToImage, err = dnld.BingWallpaper(
			viper.GetStringSlice("bing.markets"),
			viper.GetString("bing.URLResolution"),
			viper.GetInt("bing.quality"),
			viper.GetInt("bing.height"),
			viper.GetInt("bing.width"),
		)
		check(err)
	case 1:
		// Chromecast wallpaper
		pathToImage, err = dnld.ChromecastWallpaper(
			viper.GetString("chromecast.parameters"),
		)
		check(err)
	}

	// Apply wallpaper from pathToImage
	if viper.GetBool("wallpaper.applyafterdownload") {
		lib.LogInColor.Info("Applying image, this might take a while")
		err = wallpaper.SetFromFile(pathToImage)
		check(err)
		lib.LogInColor.Info("Image applied")
	}

	// Delete wallpaper
	if viper.GetBool("wallpaper.deleteafterapply") {
		err = os.Remove(pathToImage)
		check(err)
		lib.LogInColor.Info("Image deleted")
	}

	lib.LogInColor.Info("Done")
}
