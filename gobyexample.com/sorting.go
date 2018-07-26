package main

import (
	"fmt"
	"sort"
)

func main() {
	b:=[5]int{1,2,3,4,5}
	fmt.Printf("%T\n", b)
	s:=[]int{1,2,3,4,5}
	fmt.Printf("%T\n", s)
	strs:=[]string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Printf("Strings: %v; Type: %T\n", strs, strs)

	ints:=[]int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:	", ints)

	ifsorted:=sort.IntsAreSorted(ints)
	fmt.Println("Sorted:	", ifsorted)
}
