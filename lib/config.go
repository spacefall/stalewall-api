package lib

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	listOfKeys := []string{
		"ping.pingRetrySleep",
		"ping.maxPingRetries",
		"wallpaper.ApplyAfterDownload",
		"wallpaper.DeleteAfterApply",
		"chromecast.parameters",
		"bing.URLResolution",
		"bing.quality",
		"bing.height",
		"bing.width",
		"bing.markets",
	}

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

	for i := 0; i < len(listOfKeys); i++ {
		if !viper.IsSet(listOfKeys[i]) {
			LogInColor.Fatal("'" + listOfKeys[i] + "' doesn't exist in config.toml. Please set it and retry.")
		}
	}
}
