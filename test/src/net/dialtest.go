package main

import "net"

func testDial() {
	net.Dial("tcp", "192.168.0.1:8081")
	net.Dial("udp", "192.168.0.1:8081")
	net.Dial("ip4:icmp", "www.baidu.com") //使用协议名称
	//http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xml
	net.Dial("ip4:1", "www.baidu.com") //使用协议编号
}
