package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	//b := 0
	//a := 5
	//fmt.Println(a / b)
	panic(123)
}

func main() {
	tryRecover()
}
