package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1, c2 := generator(), generator()
	worker := createWorker(0)
	n := 0
	var values []int
	tm := time.After(time.Duration(18) * time.Second)
	tt := time.Tick(time.Duration(2) * time.Second)
	for {
		var realWorker chan int
		var realValue int
		if len(values) > 0 {
			realWorker = worker
			realValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case realWorker <- realValue:
			values = values[1:]
		case <-tt:
			fmt.Printf("the length of values is %d.\n", len(values))
			fmt.Println("values = ", values)
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
func generator() chan int {
	c := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

func createWorker(id int) chan int {
	c := make(chan int)
	go doWork(id, c)
	return c
}
func doWork(i int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("channel %d, received %d\n", i, n)
	}
}
