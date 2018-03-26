package main

import "fmt"

func main() {
	arr:=[...]int{0,1,2,3,4,5,6,7}
	s1:=arr[2:6]
	s2:=s1[3:5]
	fmt.Println(s2)
	s3:=append(s2,10)
	s4:=append(s3,11)
	s5:=append(s4,12)
	fmt.Println("s3, s4, s5=",s3,s4,s5)
	fmt.Println("arr = ", arr)
}
