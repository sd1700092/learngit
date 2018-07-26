package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 15})
	fmt.Println(person{age: 20, name: "John"})
	fmt.Println(person{name:"Alfred"})
	fmt.Println(&person{name:"Ann", age: 40,})
	s:=person{name:"Sean", age:50}
	fmt.Println(s.name)
	sp:=&s
	sp=&person{"Alice", 15}
	fmt.Println(sp)
	sp.age=51
	fmt.Println(sp.age)
}
