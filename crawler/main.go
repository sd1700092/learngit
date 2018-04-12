package main

import (
	//"net/http"
	//"io/ioutil"
	//"fmt"
	//"golang.org/x/text/transform"
	//"io"
	//"golang.org/x/text/encoding"
	//"bufio"
	//"golang.org/x/net/html/charset"
	//"regexp"
	"imooc.com/learngo/crawler/engine"
	"imooc.com/learngo/crawler/scheduler"
	"imooc.com/learngo/crawler/zhenai/parser"
)

func main() {
	/*resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//if resp.StatusCode == http.StatusOK {
	//	all, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%s\n", all)
	//}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	utf8Reader := transform.NewReader(resp.Body, determinEncoding(resp.Body).NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)*/
	//printCityList(all)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	//e. /*SimpleEngine{}*/ Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

/*
func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/\w+"[^>]*>[^<]+</a>`)
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
	//matches := re.FindAll(contents, -1)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//fmt.Printf("%s\n", m)
		//for _, subMatch := range m {
			//fmt.Printf("%s ", subMatch)
			fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		//}
		fmt.Println()
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}*/
