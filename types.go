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

type Data struct {
	expected       map[string]int
	msgs           map[string]string
	interval       time.Duration
	repeatDuration time.Duration
}

type Response struct {
	name string
	code int
	when time.Time
}
