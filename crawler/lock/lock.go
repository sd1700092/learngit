package main

import (
	"log"
	"strconv"
)

func main() {
	in := make(chan int)
	out := make(chan string)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				number := <-in
				out <- string(number)
			}
		}()
	}
	go func() {
		for i := 0; i < 4000; i++ {
			in <- i
		}
	}()

	for {
		result := <-out
		log.Printf("result:=%s\n", result)
		number, _ := strconv.Atoi(result)
		go func() { in <- (number + 4000) }()

	}
}
