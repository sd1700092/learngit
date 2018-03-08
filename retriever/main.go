package main

import (
	"fmt"
	"time"

	"imooc.com/learngo/retriever/mock"
	"imooc.com/learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> % T %v\n", r, r)
	fmt.Printf(" > Type switch:")
	switch v := r.(type) {
	case *mock.Aaaaa:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	////值接收者用指针赋值或者用值赋值都可以，但是用指针赋值的话Contens就打不出来了。
	//r = &mock.Aaaaa{"this is a fake imooc.com"}
	retriever := mock.Aaaaa{"this is a fake imooc.com"}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	// Type Assertion
	if Aaaaa, ok := r.(*mock.Aaaaa); ok {
		fmt.Println(Aaaaa.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
}
