package main

import (
	"fmt"
	"imooc.com/learngo/retriever/mock"
	"imooc.com/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever)  {
	fmt.Printf("% T %v\n", r, r)
	switch v:= r.(type) {
	case mock.Aaaaa:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster){
	s.Get()
	s.Post()
}

func main() {
	var r Retriever
	r = &mock.Aaaaa{"this is a fake imooc.com"} //值接收者用指针赋值或者用值赋值都可以，但是用指针赋值的话Contens就打不出来了。
	inspect(r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
	// Type Assertion
	if Aaaaa, ok := r.(mock.Aaaaa); ok{
		fmt.Println(Aaaaa.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
}
