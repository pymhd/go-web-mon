package main

import (
	"time"
	"sync"
	"sync/atomic"
)


type StatusKeeper struct {
	mu           sync.Mutex
	ResponseChan chan WebResourceResponse
	NotifyChan   chan Notification //see workers.go
	Counter	     map[string]*int32
	LastResult   interface{}
}

func (sk *StatusKeeper) handleResponses() {
	for wrr := range sk.ResponseChan {
		sk.mu.Lock()
		if wrr.StatusCode == wrr.StatusCodeExpected {
			// Good
			// we need to zeroize fail counter
			atomic.StoreInt32(sk.Counter[wrr.Name], 0)
		} else {
			atomic.AddInt32(sk.Counter[wrr.Name], 1)
			if !wrr.CodeReceived {
				sk.LastResult = "Timeout/Error"
			} else {
				sk.LastResult = wrr.StatusCode
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
                for range ticker {
                        sk.NotifyFailed()
                }
        }()

}

func (sk *StatusKeeper) NotifyFailed() {
	sk.mu.Lock()
	defer sk.mu.Unlock()
	
	for name, count := range sk.Counter {
		if *count > cfg.Web[name].Retry
		
	}
}


func NewStatusKeeper(c chan WebResourceResponse) *StatusKeeper {
	sk := new(StatusKeeper)
	sk.ResponseChan = c
	sk.NotifyChan = make(chan Notification, 50)

	return sk
}
