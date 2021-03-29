package xlimit

import (
	"log"

	"github.com/nucktwillieren/project-d/xlimit-grpc/internal"
	"github.com/nucktwillieren/project-d/xlimit-grpc/internal/xlimit"

	"google.golang.org/grpc"
)

func NewClientWithConn(address string, identity string) *XLimitClient {

	options := &internal.XLimitClientOptions{Addr: address}

	conn, err := grpc.Dial(options.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		conn.Close()
		log.Fatalf("did not connect: %v", err)
	}
	c := xlimit.NewXLimitClient(conn)

	return &c
}
