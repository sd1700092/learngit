package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n = n / 2 {
		i := n % 2
		result = strconv.Itoa(i) + result
	}
	return result
}

func main() {
	fmt.Println(convertToBin(5),
		convertToBin(13))
}
