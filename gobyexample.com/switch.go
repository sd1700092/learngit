package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Printf("Write %d as ", i)
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	day:=time.Now().Weekday()
	switch day {
	case time.Sunday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		println("It's a weekday")
	}
	t:=time.Now().Hour()
	switch  {
	case t > 12:
		println("It's after noon")
	default:
		println("It's before noon")
	}
	whatAmI:=func(i interface{}) {
		switch i.(type) {
		case bool:
			println("I am a bool")
		case int:
			println("I am an int")
		default:
			fmt.Printf("don't know type %T", i)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
