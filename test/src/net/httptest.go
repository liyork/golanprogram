package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

//net/http包

func testGet() {
	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("resp:", resp, "err:", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func testPost() {

	var reader io.Reader
	resp, err := http.Post("http://example.com/upload", "image/jpeg", reader)
	if err != nil {
		fmt.Println("resp:", resp, "err:", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("statusCode is not ok,statusCode:", resp.StatusCode)
		return
	}
}

//application/x-www-form-urlencoded
func testPostForm() {
	resp, err := http.PostForm("http://sss.com/posts", url.Values{"title": {"article title"}, "content": {"article body"}})
	if err != nil {
		fmt.Println("resp:", resp, "err:", err)
		return
	}
}

func testDo() {
	client := &http.Client{ /*CheckRedirect: redirectPolicyFunc,*/ }

	req, err := http.NewRequest("GET", "http://xcvcx.com", nil)
	req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)

	resp, err := client.Do(req)

	fmt.Println(err, resp)
}

func testTransport() {
	tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: nil}, DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	client.Get("https://sfds.com")

	//Client和Transport在执行多个 goroutine 的并发过程中都是安全的，但出于性能考虑，应当创建一次后反复使用。
}

//Head 请求方式表明只请求目标 URL 的头部信息，即 HTTP Header 而不返回 HTTP Body
//resp, err := http.Head("http://example.com/")

func main() {
	testGet()
}

//go build httptest.go
//./htttest.go
