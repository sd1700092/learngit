package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("Hello world!")
	enums()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)
	printFile("README.md")
}
