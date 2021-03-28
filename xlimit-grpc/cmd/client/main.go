package main

import (
	"context"
	"log"
	"time"

	"github.com/nucktwillieren/project-d/xlimit-grpc/internal"
	"github.com/nucktwillieren/project-d/xlimit-grpc/internal/xlimit"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	options := &internal.XLimitClientOptions{}
	envconfig.Process("XLIMIT_GRPC", options)

	conn, err := grpc.Dial(options.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := xlimit.NewXLimitClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for i := 0; i < 1000; i++ {
		r, err := c.CheckAndIncrease(ctx, &xlimit.XLimitCheckRequest{Identity: "tests", IncreaseNumber: 1})
		if err != nil && err != internal.LimitExceedError {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Greeting: %s", r.GetIdentity())
	}
}
