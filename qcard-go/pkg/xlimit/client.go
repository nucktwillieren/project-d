package xlimit

import (
	"log"

	"google.golang.org/grpc"
)

type XLimitClientOptions struct {
	Addr string
}

func NewClientWithConn(address string) *XLimitClient {

	options := &XLimitClientOptions{Addr: address}

	conn, err := grpc.Dial(options.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		conn.Close()
		log.Fatalf("did not connect: %v", err)
	}
	c := NewXLimitClient(conn)

	return &c
}
