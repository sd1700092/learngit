package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError1 string

func (e userError1) Error() string {
	return e.Message()
}

func (e userError1) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError1("path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	writer.Write(all)
	return nil
}
