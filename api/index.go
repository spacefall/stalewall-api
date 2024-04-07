package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/spaceox/stalewall/providers"
)

type settings struct {
	height int
	width  int
	crop   bool
	provs  []string
}

var providerList = []string{"b", "c", "s", "n", "u"}

func parseQueries(r *http.Request) (*settings, error) {
	config := settings{
		height: 0,
		width:  0,
		crop:   true,
		provs:  []string{"b", "c", "s", "u"},
	}
	var err error

	// parses resolution
	if r.URL.Query().Has("res") {
		res := strings.Split(r.URL.Query().Get("res"), "x")

		config.width, err = strconv.Atoi(res[0])
		if err != nil {
			return nil, err
		}

		config.height, err = strconv.Atoi(res[1])
		if err != nil {
			return nil, err
		}
	}

	// if no crop or resolution is specified, disables crop
	if r.URL.Query().Has("nc") || config.height == 0 || config.width == 0 {
		config.crop = false
	}

	// parses provider list
	if r.URL.Query().Has("p") {
		config.provs = []string{}
		provs := strings.Split(r.URL.Query().Get("p"), "")
		for _, prov := range provs {
			if !slices.Contains(providerList, prov) {
				return nil, errors.New("invalid provider list")
			}
			config.provs = append(config.provs, prov)
		}
	}
	return &config, nil
}

func getWall(config *settings) ([]byte, error) {
	// Initializing variables here to access the variables inside the switch
	var (
		imageUrl string
		provider string
		err      error
	)

	switch config.provs[rand.Intn(len(config.provs))] {
	case "b":
		// Bing wallpaper
		provider = "bing"
		imageUrl, err = providers.BingWallpaper(config.height, config.width, config.crop)
	case "c":
		// Chromecast wallpaper
		provider = "chromecast"
		imageUrl, err = providers.ChromecastWallpaper(config.height, config.width, config.crop)
	case "s":
		// Spotlight wallpaper
		provider = "spotlight"
		imageUrl, err = providers.SpotlightWallpaper(config.height, config.width)
	case "n":
		// Apod wall
		provider = "nasa"
		imageUrl, err = providers.NasaWallpaper("DEMO_KEY")
	case "u":
		// Unsplash wall
		provider = "unsplash"
		imageUrl, err = providers.UnsplashWallpaper(config.height, config.width, config.crop)
	default:
		return nil, errors.New("index out of range in switch on function getWall")
	}

	if err != nil {
		return nil, err
	}

	if imageUrl == "" {
		return nil, fmt.Errorf("provider (%v) returned empty image url", provider)
	}

	return createJSON(provider, imageUrl)

}

// creates a json, it's put here to make json modifications more streamlined
func createJSON(provider, url string) ([]byte, error) {
	data := struct {
		Provider string `json:"source"`
		URL      string `json:"url"`
	}{
		Provider: provider,
		URL:      url,
	}

	// this block does the same thing as json.marshal BUT it doesn't escape urls
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := fmt.Fprint(w, "You can't just like do that and expect me to not notice")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	config, err := parseQueries(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	wall, err := getWall(config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// identifies the page as a json
	w.Header().Set("Content-Type", "application/json")
	// allows cors
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// write the json
	_, err = w.Write(wall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
