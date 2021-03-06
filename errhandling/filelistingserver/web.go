package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"

	"imooc.com/learngo/errhandling/filelistingserver/filelisting"
	//"github.com/gpmgo/gopm/modules/log"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				//http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	//http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList)) //故意把路径写错。
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
