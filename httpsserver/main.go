package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Method: %v", request.Method)
	Printfln("URL: %v", request.URL)
	Printfln("HTTP Version: %v", request.Proto)
	Printfln("Host: %v", request.Host)
	for name, val := range request.Header {
		Printfln("Header: %v, Value: %v", name, val)
	}
	Printfln("---")
	io.WriteString(writer, sh.message)
}

func main() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
