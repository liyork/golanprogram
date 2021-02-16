package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

//使用Server接口确定之后所要实现的业务服务器的统一接口
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

//对server进行扩展
type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

//对每个连接使用新gorutine处理
func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0) //0缓冲

	go func(c chan string) { //新gorutine执行
		for {
			request := <-c //读取数据，阻塞
			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			resp := server.Handle(req.Method, req.Params)

			fmt.Println("server resp:", resp)

			b, err := json.Marshal(resp)

			c <- string(b) //写入结果
		}

		fmt.Println("Session closed")
	}(session)

	fmt.Println("A new session has been created successfully.")

	return session
}
