package main

import "fmt"

func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	pass_by_val(a)
	fmt.Printf("After pass_by_val: %d\n", a)
	pass_by_ref(&a)
	fmt.Printf("After pass_by_ref: %d\n", a)

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func pass_by_ref(a *int) {
	*a++
}

func pass_by_val(a int) {
	a++
}
