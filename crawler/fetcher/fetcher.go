package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		return nil, err
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
		//fmt.Println("Error: status code", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	utf8Reader := transform.NewReader(resp.Body, determinEncoding(resp.Body).NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determinEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}