package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

func TestErrWrapper(t *testing.T){
	tests := []struct {
		h appHandler
		code int
		message string
	} {
		{errPanic, 500, "Internal Server Error"},
	}

	for _, tt:=range tests{
		f:=errWrapper(tt.h)
		response := httptest.NewRecorder()
		request:=httptest.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
		f(response, request)
		b, _:=ioutil.ReadAll(response.Body)
		body := string(b)
		println("body="+body)
		println("tt.message="+tt.message)
		//fmt.Println("body==tt.message?" + body==tt.message)
		if response.Code!=tt.code || !strings.EqualFold(body, tt.message) {
			t.Errorf("expect(%d %s);got (%d %s)",
				tt.code, tt.message, response.Code, body)
		}
	}
}

func errPanic(writer http.ResponseWriter, request *http.Request) error{
	panic(123)
}
