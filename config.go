package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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

	//codes := make(map[string]int, 0)
	//msgs := make(map[string]string, 0)
	//for _, w := range cfg.Web {
	//	codes[w.Name] = w.ExpectedCode
	//	msgs[w.Name] = w.Msg
	//}
	//cfg.data.expected = codes
	//cfg.data.msgs = msgs

	//dur, err := time.ParseDuration(cfg.Global.Repeat)
	//must(err)
	//cfg.data.repeatDuration = dur

	//dur, err = time.ParseDuration(cfg.Global.Interval)
	//must(err)
	//cfg.data.interval = dur

	return cfg
}
