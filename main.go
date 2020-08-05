package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func getUrl(url string, timeout time.Duration, completed chan<- bool) {

	httpClient := http.Client{Timeout: timeout}
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Debugf("error with http call: %v", err)
		completed <- false
		return
	}
	log.Debugf("http response status: %+v", resp.Status)
	completed <- true
}

func main() {

	url := flag.String("url", "", "url to run the load test against")
	timeout := flag.Int("timeout", 30, "timeout for a given url")
	concurrent := flag.Int("concurrent", 1, "number of tests to run concurrently")
	duration := flag.String("duration", "30s", "duration to run tests for")
	debug := flag.Bool("debug", false, "run in debug mode with more infomation on logs")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("running in debug mode")
	}

	if *url == "" {
		flag.PrintDefaults()
		log.Fatal("flag `-url` must be passed in")
	}

	testStartedAt := time.Now()
	phrasedDuration, err := time.ParseDuration(*duration)
	if err != nil {
		log.Fatal("please enter a valid duration for duration. e.g: 20s, 20m")
	}
	testEndAt := testStartedAt.Add(phrasedDuration)

	successfulTests := 0
	failedTests := 0
	completed := make(chan bool)

	log.Infof("load testing: %v", *url)

	for i := 1; i <= *concurrent; i++ {
		go getUrl(*url, time.Duration(*timeout)*time.Second, completed)
	}

	for success := range completed {

		if success {
			successfulTests += 1
		} else {
			failedTests += 1
		}
		log.Infof("test completed: %v", success)

		timeNow := time.Now()
		if testEndAt.Before(timeNow) {
			break
		}
		go getUrl(*url, time.Duration(*timeout)*time.Second, completed)
	}

	fmt.Println("tested url: ", *url)
	fmt.Println("test started at: ", testStartedAt)
	fmt.Println("test ended at: ", testEndAt)
	fmt.Println("successful tests: ", successfulTests)
	fmt.Println("failed tests: ", failedTests)
}
