package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("work", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("work", id, "finish job", j)
		results <- j * 2
	}
	//close(results)
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int)
	for i := 0; i < 3; i++ {
		go work(i, jobs, results)
	}
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//close(results)
	for a:=1;a<=5;a++{
		fmt.Println("result = ", <-results)
	}
	//for result := range results {
	//	fmt.Println("result = ", result)
	//}
}
