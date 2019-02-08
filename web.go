package main

import (
	"net/http"
	"time"
)

const (
	Timeout = 11
)

func getResponseCode(url string) int {
	codes := make(chan int)

	go func(u string) {
		resp, err := http.Head(u)
		if err != nil {
			codes <- 0
		}
		defer resp.Body.Close()
		codes <- resp.StatusCode
	}(url)

	select {
	case <-time.After(5 * time.Second):
		return Timeout
	case code := <-codes:
		return code
	}
}
