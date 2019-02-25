package main

import (
	"time"
)

/*
type MessageRenderer struct {
	URL     string
	Name    string
	ExpCode int
	RcvCode int
}
*/

type RuntimeConfig struct {
	workers  int
	interval time.Duration
	repeat   time.Duration
	messages map[int]string
	codes    map[int]int
	objects  []string
}

func CreateRuntimeConfig(c *Config) *RuntimeConfig {
	rc := new(RuntimeConfig)

	rc.messages = make(map[int]string, 0)
	rc.codes = make(map[int]int, 0)
	rc.objects = make([]string, len(c.Web))

	dur, err := time.ParseDuration(c.Global.Repeat)
	must(err)
	rc.repeat = dur

	dur, err = time.ParseDuration(c.Global.Interval)
	must(err)
	rc.interval = dur

	for i, w := range c.Web {
		rc.codes[i] = w.ExpectedCode
		rc.messages[i] = w.Msg
		rc.objects[i] = w.URL
	}
	return rc
}

/*
func mustRenderMessage(s string, i interface{}) string {
	var b bytes.Buffer

	t := template.Must(template.New("msg").Parse(s))
	t.Execute(b, i)

	return b.String()
}
*/
