package main

import (
	"github.com/gorilla/mux"
	"github.com/shitpostingio/analysis-commons/decoder"
	"github.com/shitpostingio/analysis-commons/handler"
	health_check "github.com/shitpostingio/analysis-commons/health-check"
	"github.com/shitpostingio/fingerprint-microservice/fingerprinting"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	imageSizeKey   = "FP_MAX_IMAGE_SIZE"
	videoSizeKey   = "FP_MAX_VIDEO_SIZE"
	bindAddressKey = "FP_BIND_ADDRESS"
)

var (
	maxImageSize int64 = 10 << 20 // 10MB
	maxVideoSize int64 = 20 << 20 // 20MB
	bindAddress        = "localhost:10000"
	r            *mux.Router
)

func main() {

	r.HandleFunc("/fingerprinting/image", handleImage).Methods("POST")
	r.HandleFunc("/fingerprinting/video", handleVideo).Methods("POST")
	r.HandleFunc("/healthy", health_check.ConfirmServiceHealth).Methods("GET")
	log.Println("Fingerprinting server powered on!")
	log.Fatal(http.ListenAndServe(bindAddress, r))

}

func handleImage(w http.ResponseWriter, r *http.Request) {
	handler.Handle(w, r, maxImageSize, &decoder.ImageDecoder{}, fingerprinting.GetFingerprint)
}

func handleVideo(w http.ResponseWriter, r *http.Request) {
	handler.Handle(w, r, maxVideoSize, &decoder.VideoDecoder{}, fingerprinting.GetFingerprint)
}
