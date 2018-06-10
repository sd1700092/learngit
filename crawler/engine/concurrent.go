package engine

import (
	"log"
	"imooc.com/learngo/crawler/model"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	//itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile #%d: %v\n", profileCount, item)
				profileCount++
			}
			log.Printf("Got item #%d: %v", 0, item)
		}
		for _, request := range result.Requests {
			if (isDuplicate(request.Url)) {
				//log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
