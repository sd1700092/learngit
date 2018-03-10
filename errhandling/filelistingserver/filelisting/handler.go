package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func(e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):] //request.URL.Path="/abc" 直接在这步就报错了
	file, err := os.Open(path)
	if (err != nil) {
		//http.Error(writer, err.Error(), http.StatusInternalServerError) //这句话会在浏览器上打印出出错信息
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
		return err
	}
	writer.Write(all)
	return nil
}

func main() {
}
