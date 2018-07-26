package main

import (
	"crypto/sha1"
	"fmt"
	"encoding/hex"
)

func main() {
	h:=sha1.New()
	s:="sha1 this string"
	h.Write([]byte(s))
	bs:=h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Printf(hex.EncodeToString(bs))
}
