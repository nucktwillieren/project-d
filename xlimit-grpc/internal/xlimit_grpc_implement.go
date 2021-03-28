package internal

import (
	context "context"
	"log"
	"net"

	"github.com/kelseyhightower/envconfig"
	"github.com/nucktwillieren/project-d/xlimit-grpc/internal/xlimit"
	grpc "google.golang.org/grpc"
)

type xLimitService struct {
	xlimit.UnimplementedXLimitServer
	Addr  string
	layer *XlimitRedisLayer
}

type XLimitClientOptions struct {
	Addr string
}

func (x *xLimitService) CheckAndIncrease(ctx context.Context, in *xlimit.XLimitCheckRequest) (*xlimit.XLimitCheckReply, error) {
	out := &xlimit.XLimitCheckReply{}
	_, out, err := x.layer.CheckAndIncrease(ctx, in, out)
	log.Println(out)
	return out, err
}

func (x *xLimitService) Reset(ctx context.Context, in *xlimit.XLimitResetRequest) (*xlimit.XLimitCheckReply, error) {
	log.Printf("Identity: %v", in.GetIdentity())
	return &xlimit.XLimitCheckReply{}, nil
}

func NewXLimitService(layer *XlimitRedisLayer) *xLimitService {
	var service xLimitService

	if err := envconfig.Process("XLIMIT_GRPC", &service); err != nil {
		log.Fatalf("failed to create xLimitService: %v", err)
	}
	log.Println(service.Addr)

	lis, err := net.Listen("tcp", service.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service.layer = layer

	s := grpc.NewServer()
	xlimit.RegisterXLimitServer(s, &service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return &service
}