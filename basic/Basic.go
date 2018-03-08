package main

import (
	"fmt"
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
