package main 

import (
        "time"
        "net/http"
)

type Notification struct {
        To		int
        Message		string
}

type WebResourceResponse struct {
        Name		string
        StatusCode	int
        When		time.Time
        CodeReceived	bool
}


func runWorkers(c  int) {
        i := make(chan WebResource)
        o := make(chan WebResourceResponse)
        
        //deploy new status keeper with response channel that will be handeled by it's goroutines
        sk := NewStatusKeeper(o)
        
        for i < c; i++ {
                go runWebResourceMonitor(i, o)
        }
        //start tlg notificators = watchers / 10 ( for instance if all resources will go down simultaneously we will notify users in few threads)
        var nc := 1
        if x:= c/10; x >0  {
                nc = x
        } 
        //start tlg notificators
        for i < nc; i++ {
                go runNotificator(sk.NotifyChan)
        }
}


func runWebResourceMonitor(in chan WebResource, out chan WebResourceResponse) {
        for wr := range in {
                code, ok := getResponseCode(wr.URL, wr.Timeout)
                now := time.Now()
                out <- WebResourceResponse{Name: wr.Name, StatusCode: code, When: now, CodeReceived: ok}
        }
}



func getResponseCode(url string, timeout time.Duration) (int, bool) {
	codes := make(chan int, 1)

	go func(u string) {
		resp, err := http.Head(u)
		if err != nil {
			codes <- 0
			return 
		}
		defer resp.Body.Close()
		codes <- resp.StatusCode
	}(url)

	select {
	case <- time.After(timeout):
		return (0, false)
	case code := <-codes:
		return (code, true)
	}
}


func runNotificator(in chan Notification) {
                
}
