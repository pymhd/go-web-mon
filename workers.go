package main

import (
	"fmt"
	"time"

	log "github.com/pymhd/go-logging"
	bot "github.com/pymhd/tlgrm-bot"
)

func DispatchWorkers(n int) (chan WebResource, chan Response) {
	log.Infof("Dispatching %d workers\n", n)
	outputChannel := make(chan Response, n)
	inputChannel := make(chan WebResource, n*100)

	for i := 0; i < n; i++ {
		go RunWorker(i, inputChannel, outputChannel)
	}
	return inputChannel, outputChannel
}

func DispatchNotificators(n int, c chan Response) {
	log.Infof("Dispatching %d notificators\n", n)
	for i := 0; i < n; i++ {
		go RunNotificator(i, c)
	}
}

func RunWorker(i int, in chan WebResource, out chan Response) {
	log.Infof("Worker - %d started\n", i)
	for wr := range in {
		code := getResponseCode(wr.URL)
		now := time.Now()
		r := Response{
			name: wr.Name,
			code: code,
			when: now,
		}
		log.Infof("Worker - %d: %s returned %d status code\n", i, wr.URL, code)
		out <- r
	}
}

func RunNotificator(i int, c chan Response) {
	log.Infof("Notificator - %d started\n", i)
	for resp := range c {
		if needToNotify(resp) {
			go SendTlgMessages(resp)
		}
	}
}

func needToNotify(r Response) bool {
	st := keeper.getState(r.name)
	if expCode := cfg.data.expected[r.name]; r.code != expCode {
		log.Infof("Expected status code for %s did not match (expected: %d, Got: %d)\n", r.name, expCode, r.code)
		log.Infof("Alarm status for %s is %t\n", r.name, st.isActive)
		if !st.isActive || st.isActive && time.Since(st.when) > cfg.data.repeatDuration {
			log.Infof("Users need to be notified about %s\n", r.name)
			keeper.setState(r.name, true, time.Now())
			log.Infoln("Set alarm state as true, updated alarm timestamp")
			return true
		} else {
			log.Infof("Wont invoke notification for %s\n", r.name)
			return false
		}
	} else {
		if st.isActive {
			log.Infof("Not implemented restore handler\n")
			keeper.setState(r.name, false, time.Now())
		}
		return false
	}
}

func SendTlgMessages(r Response) {
	msg := cfg.data.msgs[r.name]
	expCode := cfg.data.expected[r.name]
	for _, chatId := range cfg.Tlg.Chats {
		m := fmt.Sprintf(msg, r.name, expCode, r.code)
		msgId, err := bot.SendTextMessage(cfg.Tlg.Token, chatId, m, 0)
		if err != nil {
			log.Errorln(err)
		}
		log.Infof("Sent tlg message, id is %d\n", msgId)
	}
}
