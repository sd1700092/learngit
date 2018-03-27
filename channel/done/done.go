package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, wg *sync.WaitGroup) {
	/*	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c\n", id, n)
	}*/
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		//go func() { done <- true }()
		wg.Done()
	}
}

type worker struct {
	in chan int
	//done chan bool
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{make(chan int), wg}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	//var c chan int // create a channel of int where c == nil
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	wg.Wait()
	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}

func main() {
	chanDemo()

}
