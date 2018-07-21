package main

import "fmt"

func main() {
	m:=make(map[string]int)
	m["k1"]=7
	m["k2"]=13
	fmt.Println(m)
	fmt.Println(m["k1"])
	fmt.Println(len(m))
	delete(m, "k2")
	fmt.Println(m)
	_, prs:=m["k1"]
	fmt.Println(prs)
	m=map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)
	m=make(map[string]int)
	m=map[string]int{"a": 1, "b": 2}
}
