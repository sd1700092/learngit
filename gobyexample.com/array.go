package main

import "fmt"

func main() {
	var emp [5]int
	fmt.Println("emp:", emp)
	emp[4] = 100
	fmt.Println("set:", emp)
	fmt.Println("get:", emp)
	fmt.Println("len:", len(emp))
	b:=[5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:", twoD)

	arr:=make([][]int, 3)
	for i:=0;i<3;i++{
		innerLen:=i+1
		arr[i]=make([]int, innerLen)
		for j:=0;j<innerLen;j++{
			arr[i][j]=i+j
		}
	}
	fmt.Println("2d:", arr)
}
