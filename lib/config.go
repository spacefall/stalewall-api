package lib

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Default config
	// Ping
	viper.SetDefault("ping.pingRetrySleep", "5s")
	viper.SetDefault("ping.maxPingRetries", 5)
	// Wallpaper application
	viper.SetDefault("wallpaper.ApplyAfterDownload", true)
	viper.SetDefault("wallpaper.DeleteAfterApply", true)
	// Chromecast
	viper.SetDefault("chromecast.parameters", "h0-w0")
	// Bing
	viper.SetDefault("bing.URLResolution", "UHD")
	viper.SetDefault("bing.quality", 100)
	viper.SetDefault("bing.height", 0)
	viper.SetDefault("bing.width", 0)
	viper.SetDefault("bing.markets", []string{"en-US", "ja-JP", "en-AU", "en-UK", "de-DE", "en-NZ", "en-CA"})

	// Config name, extension and path
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// Check if config is loading fine
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
