package client

import (
	"bytes"
	"encoding/json"
	"github.com/shitpostingio/analysis-commons/structs"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// PerformRequest performs a request to the fingerprinting service.
func PerformRequest(file io.Reader, fileName, endpoint string) (data structs.FingerprintResponse, errorString string) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		errorString = err.Error()
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		errorString = err.Error()
		return
	}

	err = writer.Close()
	if err != nil {
		errorString = err.Error()
		return
	}

	log.Debugln("Fpclient.PerformRequest: preparing fingerprinting request for file ", fileName, " to endpoint ", endpoint)
	request, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		errorString = err.Error()
		return
	}

	// We want to send data with the multipart/form-data Content-Type
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{Timeout: time.Second * 30}
	response, err := client.Do(request)
	if err != nil {
		errorString = err.Error()
		return
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Println("Fpclient.PerformRequest: unable to close response body", err)
		}
	}()

	bodyResult, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Debugln("Fpclient.PerformRequest: error in result: ", string(bodyResult))
		errorString = err.Error()
		return
	}

	var ar structs.Analysis
	err = json.Unmarshal(bodyResult, &ar)
	if err != nil {
		errorString = err.Error()
		log.Println("PerformRequest: error while unmarshaling ", err, " body result was: ", string(bodyResult))
		return
	}

	return ar.Fingerprint, ar.FingerprintErrorString

}
