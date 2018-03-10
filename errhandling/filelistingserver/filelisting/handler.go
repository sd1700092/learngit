package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error{
	path := request.URL.Path[len("/list/"):]
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
