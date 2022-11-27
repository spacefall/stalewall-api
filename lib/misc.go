package lib

import (
	"os"
	"path/filepath"

	"github.com/withmandala/go-log"
)

var LogInColor = log.New(os.Stdout)

// Returns full path
func GetFinalPath(filename string) (string, error) {

	workdir, err := os.Getwd()
	pathToImage := filepath.Join(workdir, filename)
	if err != nil {
		return "", err
	}
	LogInColor.Info("Path to image:", pathToImage)
	return pathToImage, nil
}
