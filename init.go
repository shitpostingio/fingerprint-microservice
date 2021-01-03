package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func init() {
	setEnvVars()
	r = mux.NewRouter()
}

func setEnvVars() {

	mis := os.Getenv(imageSizeKey)
	if mis != "" {
		p, err := strconv.ParseInt(mis, 10, 64)
		if p > 0 && err == nil {
			maxImageSize = p
			log.Println("Found environment variable! New max image size: ", p)
		}
	}

	mvs := os.Getenv(videoSizeKey)
	if mvs != "" {
		p, err := strconv.ParseInt(mvs, 10, 64)
		if p > 0 && err == nil {
			maxVideoSize = p
			log.Println("Found environment variable! New max video size: ", p)
		}
	}

	add := os.Getenv(bindAddressKey)
	if add != "" {
		bindAddress = add
		log.Println("Found environment variable! New bind address: ", add)
	}

}
