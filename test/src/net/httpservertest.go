package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func testListenAndServe() {
	http.Handle("/foo", nil /*fooHandler*/)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	//第二个参数表示服务端处理程序，通常为空，这意味着服务端调用 http.DefaultServeMux 进行处理，而服务端编写的业务逻
	//辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testServe() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        nil /*myHandler*/,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func testListenAndServeTLS() {
	http.Handle("/foo", nil /*fooHandler*/)
	//只处理HTTPS请求，服务器上必须存在包含证书和与之匹配的私钥的相关文件，比如certFile对应SSL证书
	//文件存放路径， keyFile对应证书私钥文件路径。如果证书是由证书颁发机构签署的， certFile
	//参数指定的路径必须是存放在服务器上的经由CA认证过的SSL证书
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}

func testServeTLS() {
	ss := &http.Server{
		Addr:           ":10443",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))
}
