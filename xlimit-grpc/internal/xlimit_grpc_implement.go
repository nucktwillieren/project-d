package internal

import (
	context "context"
	"log"
	"net"

	"github.com/kelseyhightower/envconfig"
	"github.com/nucktwillieren/project-d/xlimit-grpc/pkg/xlimit"
	grpc "google.golang.org/grpc"
)

type xLimitService struct {
	xlimit.UnimplementedXLimitServer
	Addr  string
	layer *XlimitRedisLayer // I think, someday, we will use redis and the other databases for this service simultaneously, so using this pattern.
}

func (x *xLimitService) CheckAndIncrease(ctx context.Context, in *xlimit.XLimitCheckRequest) (*xlimit.XLimitCheckReply, error) {
	out := &xlimit.XLimitCheckReply{}
	_, out, err := x.layer.CheckAndIncrease(ctx, in, out)
	log.Println(out)
	return out, err
}

func (x *xLimitService) Get(ctx context.Context, in *xlimit.XLimitGetRequest) (out *xlimit.XLimitGetReply, err error) {
	out = &xlimit.XLimitGetReply{}
	log.Println(in.GetKeysOnly())
	if in.GetKeysOnly() {
		out, err = x.layer.GetKeysOnly(ctx, in, out)
	} else {
		out, err = x.layer.Get(ctx, in, out)
	}
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
