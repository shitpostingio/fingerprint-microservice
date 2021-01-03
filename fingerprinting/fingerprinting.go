package fingerprinting

import (
	"github.com/AlessandroPomponio/hsv/histogram"
	"github.com/corona10/goimagehash"
	"github.com/shitpostingio/analysis-commons/decoder"
	"github.com/shitpostingio/analysis-commons/structs"
	log "github.com/sirupsen/logrus"
	"image"
	"io"
)

// GetFingerprint fingerprints a file.
func GetFingerprint(extension string, reader io.Reader, f decoder.MediaDecoder) *structs.Analysis {

	img, err := f.Decode(extension, reader)
	if err != nil {
		log.Printf("GetFingerprint: unable to decode file with extension %s: %s\n", extension, err)
		return &structs.Analysis{FingerprintErrorString: err.Error()}
	}

	fr, err := fingerprint(img)
	if err != nil {
		return &structs.Analysis{Fingerprint: fr, FingerprintErrorString: err.Error()}
	}

	return &structs.Analysis{Fingerprint: fr}

}

func fingerprint(img image.Image) (fr structs.FingerprintResponse, err error) {

	hash, err := goimagehash.PerceptionHash(img)
	if err != nil {
		log.Println("GetFingerprint.PerceptionHash:", err)
		return
	}
	fr.PHash = hash.ToString()

	fr.Histogram = histogram.With32BinsConcurrent(img, histogram.RoundClosest)
	return

}
