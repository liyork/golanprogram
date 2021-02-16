package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	//解析地址和端口
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError2(err)

	//建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError2(err)

	conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError2(err)

	result, err := ioutil.ReadAll(conn)
	checkError2(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError2(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

//cd net
//go build simplehttp2.go
//./simplehttp2 www.baidu.com:80

//net包
//验证IP地址有效性：net.ParseIP()
//创建子网掩码：func IPv4Mask(a, b, c, d byte) IPMask
//获取默认子网掩码：func (ip IP) DefaultMask() IPMask
//根据域名查找IP
//func ResolveIPAddr(net, addr string) (*IPAddr, error)
//func LookupHost(name string) (cname string, addrs []string, err error)；
