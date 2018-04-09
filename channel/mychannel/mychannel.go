package main

import "fmt"

func main() {
	chanDemo()
}

type Worker1 struct {
	in   chan int
	done chan bool
}

func createWork(id int) Worker1 {
	w := Worker1{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func doWork(i int, in chan int, done chan bool) {
	for n := range in {
		fmt.Printf("channel %d, received %c\n", i, n)
		go func() { done <- true }()
	}
}

func chanDemo() {
	var workers [10]Worker1
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}
}
