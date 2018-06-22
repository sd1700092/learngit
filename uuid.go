package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(err)
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	fmt.Printf("%s\n", string(uuid))
	for _, b := range uuid {
		fmt.Printf("%x, ", b)
	}
	fmt.Println()
	fmt.Printf("%X\n", uuid[:])
	fmt.Printf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
