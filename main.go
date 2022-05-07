package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	h := handler{}
	http.Handle("/", &h)
	http.ListenAndServe(":9539", nil)
}

type handler struct {

}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("receive a request")
	fmt.Println(request.RequestURI)
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("fail to read from request body")
		return
	}
	fmt.Println(string(data))
}
