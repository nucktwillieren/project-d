package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/nucktwillieren/project-d/xlimit-grpc/internal"
	"github.com/nucktwillieren/project-d/xlimit-grpc/internal/xlimit"

	"google.golang.org/grpc"
)

var address string
var identity string
var times int

func main() {
	// Set up a connection to the server.
	flag.StringVar(&address, "a", "127.0.0.1:50031", "server address(ip:port) (default=127.0.0.1:50031)")
	flag.StringVar(&identity, "i", "test", "the identity you wanna test")
	flag.IntVar(&times, "n", 1001, "send n times request to server (default=1001)")
	flag.Parse()

	log.Printf("Test(%v:%v): %v times", address, identity, times)

	options := &internal.XLimitClientOptions{Addr: address}

	conn, err := grpc.Dial(options.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := xlimit.NewXLimitClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for i := 0; i < times; i++ {
		r, err := c.CheckAndIncrease(ctx, &xlimit.XLimitCheckRequest{Identity: identity, IncreaseNumber: 1})
		if err != nil && err != internal.LimitExceedError {
			log.Fatalf("could not check: %v", err)
		}

		log.Printf("Check %s: %v(Remain:%v)(Reset:%v)", r.GetIdentity(), r.GetIsAllowed(), r.GetCountRemaining(), r.GetTimeleft())
	}
}
