package main

import "fmt"

func main() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	copy(s2, s1)
	fmt.Println(s1, s2, s3)
}
