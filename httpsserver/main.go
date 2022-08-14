package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//Printfln("Method: %v", request.Method)
	//Printfln("URL: %v", request.URL)
	//Printfln("HTTP Version: %v", request.Proto)
	//Printfln("Host: %v", request.Host)
	//for name, val := range request.Header {
	//	Printfln("Header: %v, Value: %v", name, val)
	//}
	//Printfln("---")
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
	//switch request.URL.Path {
	//case "/fevicon.ico":
	//	http.NotFound(writer, request)
	//case "/message":
	//	io.WriteString(writer, sh.message)
	//default:
	//	http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	//}
}

func main() {
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.cer", "certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
