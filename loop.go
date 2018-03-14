package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"io/ioutil"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n = n / 2 {
		i := n % 2
		result = strconv.Itoa(i) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err!=nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

const filename = "README.md"

func main() {
	fmt.Println(convertToBin(5),
		convertToBin(13))
	printFile("README.md")
	if contents, err := ioutil.ReadFile(filename); err!=nil{
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}

}
