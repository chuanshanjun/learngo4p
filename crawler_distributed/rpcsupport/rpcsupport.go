package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要开Serve那就开serveRpc
func ServeRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	log.Printf("Listening port %s", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	// 虽然这里到不了，但还是要写一句return
	return nil
}

// 要开client就连client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
