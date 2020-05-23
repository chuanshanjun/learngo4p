package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	rpcdemo "chuanshan.github.com/learngo4p/lang/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	// 此处result传递的是指针，那么方法在操作的时候改的是指针地址上的值
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	// 所以返回的时候咱直接还是读变量值，因为它是一个东西，而不是值拷贝
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
