package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//建立连接
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}
	var reply int
	//同步
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith:%d*%d=%d", args.A, args.B, reply)

	//异步
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	replayCall := <-divCall.Done

	fmt.Printf("replayCall:", replayCall)
}
