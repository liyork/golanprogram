package main

import (
	"io"
	"log"
	"net/http"
)

//helloHandler()方法是http.HandlerFunc类型的实例
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!")
}

//go run helloweb.go
//浏览器访问：http://localhost:8080/hello
func main() {
	http.HandleFunc("/hello", helloHandler) //映射
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
