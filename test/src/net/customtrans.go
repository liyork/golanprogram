package main

import "net/http"

type OurCustomTransport struct {
	Transport http.RoundTripper
}

func (t *OurCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

//实现了http.RoundTripper 接口的代码通常需要在多个 goroutine中并发执行，因此必须确保实现代码的线程安全性
func (t *OurCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	//处理业务
	//发起http请求
	//添加一些域到req.Header中
	return t.transport().RoundTrip(req)
}

func (t *OurCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func main() {
	t := &OurCustomTransport{}
	client := t.Client()
	resp, err := client.Get("http://xcvcx.com")
}
