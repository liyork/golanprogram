package ipc

import "testing"

//实现Server方法
type EchoServer struct {
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"200", "ECHO:" + method + ":" + params}
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("method1", "p1")
	resp2, _ := client2.Call("method2", "p2")

	if resp1.Body != "ECHO:method1:p1" || resp2.Body != "ECHO:method2:p2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
