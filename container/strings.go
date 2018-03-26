package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网！"
	fmt.Println(len(s))
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)

	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch) //(0 59) (1 65) (2 73) (3 6211) (6 7231) (9 6155) (12 8BFE) (15 7F51) (18 FF01)
		// 这里的ch就是一个rune
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch,size :=utf8.DecodeRune(bytes)
		bytes=bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	//ch, size := utf8.DecodeRune(bytes)
	for i, ch:=range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
