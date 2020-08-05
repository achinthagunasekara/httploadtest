package main

import (
	log "github.com/sirupsen/logrus"
	"http"
)

func main() {
    log.Info("hello world")
	httpClient := http.Client{Timeout: timeout}
	resp, err := httpClient.Get("https://54dv9b2pi7.execute-api.us-east-1.amazonaws.com")
	if err != nil {
		log.Errorf("error with http call to: %v, details: %v", url, err)
	}
	log.Info(resp)
}
