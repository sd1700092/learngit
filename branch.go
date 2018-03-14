package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score > 100 || score < 0:
		//panic(fmt.Sprintf("Wrong score: %d\n", score))
		panic(score)
	case score < 60:
		g = "F"
	case score < 70:
		g = "E"
	case score < 80:
		g = "D"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "README.md"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
		fmt.Println(string(contents))
	}
	fmt.Println(grade(74))
	fmt.Println(grade(-1))
}
