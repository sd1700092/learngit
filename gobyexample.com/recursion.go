package main

import "fmt"

func main() {
	fmt.Println(fact(7))
}

func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}
