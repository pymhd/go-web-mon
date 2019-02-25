package main

import (
	"time"
)

type WebResource struct {
	Name         string `yaml:"name"`
	URL          string `yaml:"url"`
	ExpectedCode int    `yaml:"expectedCode"`
	Msg          string `yaml:"msg"`
}


type Response struct {
	name string
	code int
	when time.Time
}
