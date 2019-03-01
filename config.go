package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type WebResource struct {
	Name         string                 `yaml:"name"`
	URL          string                 `yaml:"url"`
	ExpectedCode int                    `yaml:"expectedCode"`
	Msg          string                 `yaml:"msg"`
	Interval     string                 `yaml:"interval"`
	Retry        int                    `yaml:"retry"`
	Timeout      string                 `yaml:"timeout"`
	Labels       map[string]interface{} `yaml:"labels"`
}

type Config struct {
	Global struct {
		Interval string `yaml:"interval"`
		Workers  int    `yaml:"workers"`
		Repeat   string `yaml:"repeat"`
	} `yaml:"global"`
	Tlg struct {
		Token string `yaml:"token"`
		Chats []int  `yaml:"chats"`
	} `yaml:"tlg"`
	Web []WebResource `yaml:"web"`
}

func ParseConfig(filename string) *Config {
	cfg := new(Config)
	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(fb, cfg); err != nil {
		panic(err)
	}

	return cfg
}
