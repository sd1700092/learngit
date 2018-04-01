package main

import (
	"regexp"
	"fmt"
)

//const text = "My email is ccmouse@gmail.com@abc.com"
const text = `my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is     kkk@qq.com
email2 is ddd@abc.com.cn
`

func main() {
	//re, err := regexp.Compile("ccmouse@gmail.com")
	//re := regexp.MustCompile("ccmouse1@gmail.com")
	//re := regexp.MustCompile(`\w+@\w+\.\w+`)
	//match := re.FindString(text)
	re := regexp.MustCompile(`(\w+)@([\w]+)(\.[\w.]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m:=range match {
		fmt.Println(m)
	}
	fmt.Println(match)
}
