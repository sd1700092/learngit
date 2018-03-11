package main

import (
	"fmt"
	"math"
)

func enums() {
	const (
		cpp = iota //自增序列
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func triangle(){
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

func test(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v, "is string")
	case int:
		fmt.Println(v, "is int")
	}
}



func main() {
	fmt.Println("Hello world!")
	enums()
	test("hello world")
	test(1)
}
