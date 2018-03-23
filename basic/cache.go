package main

import "fmt"

func printArray(arr [5]int){
	arr[0] = 100
	for i, v := range arr{
		fmt.Println(i, v)
	}
	//arr[0] = 199
}

func main() {
	cache := make(map[string]string)
	cache["name"] = "ccmouse"
	var array1 [5]int
	array2 := [3]int{1, 2, 3}
	array3 := [...]int{1, 2, 3, 4, 5}
	//var grid [4][5] bool
	fmt.Println(array1, array2, array3)
	for i := range array3 {
		fmt.Println(array3[i])
	}
	printArray(array3)
}
