package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/nucktwillieren/project-d/xlimit-grpc/pkg/xlimit"

	"google.golang.org/grpc"
)

var address string
var identity string
var keys bool

func main() {
	// Set up a connection to the server.
	flag.StringVar(&address, "a", "127.0.0.1:50031", "server address(ip:port) (default=127.0.0.1:50031)")
	flag.StringVar(&identity, "i", "test", "the identity you wanna test")
	flag.BoolVar(&keys, "k", false, "-k to get keys and values")
	flag.Parse()

	log.Printf("Test(%v:%v)", address, identity)

	options := &xlimit.XLimitClientOptions{Addr: address}

	conn, err := grpc.Dial(options.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := xlimit.NewXLimitClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := c.Get(ctx, &xlimit.XLimitGetRequest{Identity: identity, KeysOnly: keys})
	if err != nil && err != xlimit.LimitExceedError {
		log.Fatalf("could not check: %v", err)
	}
	for _, v := range r.Results {
		log.Printf("Get: %v", v)
	}
}
