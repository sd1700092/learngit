package main

import (
	"fmt"
	"sync"
)

type worker1 struct {
	in   chan int
	done func()
}

func main() {
	chanDemo1()
}
func chanDemo1() {
	var workers [10]worker1
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		wg.Add(1)
		worker.in <- 'A' + i
	}
	wg.Wait()
}
func createWorker(id int, wg *sync.WaitGroup) worker1 {
	w := worker1{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork1(id, w)
	return w
}
func doWork1(id int, w worker1) {
	for n := range w.in {
		fmt.Printf("channel %d, received %c\n", id, n)
		w.done()
	}
}
