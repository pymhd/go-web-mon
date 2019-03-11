package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type StatusKeeper struct {
	mu           sync.Mutex
	ResponseChan chan WebResourceResponse
	NotifyChan   chan Notification //see workers.go
	Counter      map[string]*int32
	LastResult   map[string]interface{}
}

func (sk *StatusKeeper) handleResponses() {
	for wrr := range sk.ResponseChan {
		sk.mu.Lock()
		if wrr.StatusCode == wrr.ExpectedCode {
			// Good
			// we need to zeroize fail counter
			atomic.StoreInt32(sk.Counter[wrr.Name], 0)
		} else {
			atomic.AddInt32(sk.Counter[wrr.Name], 1)
			if !wrr.CodeReceived {
				sk.LastResult[wrr.Name] = "Timeout/Error"
			} else {
				sk.LastResult[wrr.Name] = wrr.StatusCode
			}
		}
		sk.mu.Unlock()
	}
}

func (sk *StatusKeeper) Run() {
	//run chan receiver in background
	// it will get, parse and count responses for web resource
	go sk.handleResponses()

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			sk.NotifyFailed()
		}
	}()
}


func (sk *StatusKeeper) NotifyFailed() {
	sk.mu.Lock()
	defer sk.mu.Unlock()

	for name, count := range sk.Counter {
		if *count > int32(cfg.Web[name].Retry) {
		        sk.NotifyChan <- Notification{10, "test"}
		}
	}
}

func NewStatusKeeper(c chan WebResourceResponse) *StatusKeeper {
	sk := new(StatusKeeper)
	sk.ResponseChan = c
	sk.NotifyChan = make(chan Notification, 50)

	return sk
}

