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

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
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

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	a, b, c, s := 3, 4, true, "def"
	//var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func main() {
	fmt.Println("Hello world!")
	enums()
	test("hello world")
	test(1)
	variableInitialValue()
	variableTypeDeduction()
	consts()
}
