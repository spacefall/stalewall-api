package lib

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	listOfRequiredKeys := []string{
		// Ping
		"ping.pingRetrySleep",
		"ping.maxPingRetries",
		// Wallpaper
		"wallpaper.ApplyAfterDownload",
		"wallpaper.DeleteAfterApply",
		"wallpaper.QuietMode",
		// Chromecast
		"chromecast.parameters",
		// Bing
		"bing.URLResolution",
		"bing.quality",
		"bing.markets",
		// Spotlight
		"spotlight.locales",
		"spotlight.portrait",
	}
	/*
		These keys are not required:
			bing.height
			bing.width
	*/

	// Config name, extension and path
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			LogInColor.Fatal("config.toml doesn't exist. Please create one and retry.")
		} else {
			panic(err)
		}
	}

	for i := 0; i < len(listOfRequiredKeys); i++ {
		if !viper.IsSet(listOfRequiredKeys[i]) {
			LogInColor.Fatal("'" + listOfRequiredKeys[i] + "' doesn't exist in config.toml. Please set it and retry.")
		}
	}
}
