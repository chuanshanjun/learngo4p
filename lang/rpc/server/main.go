package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "chuanshan.github.com/learngo4p/lang/rpc"
)

func main() {
	// 1 rpc注册
	rpc.Register(rpcdemo.DemoService{})
	// 2 监听端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		// 如果有多个conn 可以直接用go开出多个goroutine
		go jsonrpc.ServeConn(conn)
	}
}
