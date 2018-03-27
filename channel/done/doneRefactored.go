package main

import (
	"fmt"
	"sync"
)

func doWgWork(id int, w WgWorker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type WgWorker struct {
	in   chan int
	done func()
}

func createWgWorker(id int, wg *sync.WaitGroup) WgWorker {
	w := WgWorker{in: make(chan int), done: func() { wg.Done() }}
	go doWgWork(id, w)
	return w
}

func wgChanDemo() {
	var wg sync.WaitGroup
	var workers [10]WgWorker
	for i := 0; i < 10; i++ {
		workers[i] = createWgWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	wgChanDemo()
}
