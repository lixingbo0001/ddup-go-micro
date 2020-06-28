package grpc

import (
	"github.com/lixingbo0001/ddup-go-util/util"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
)

func Handle(registryAddress string, serviceName string, serviceAddress string) *grpcService {
	reg := consul.NewRegistry(
		registry.Addrs(registryAddress),
	)

	rpcService := grpc.NewService(
		micro.Address(serviceAddress),
		micro.Name(serviceName),
		micro.Registry(reg),
	)

	return &grpcService{
		service: rpcService,
	}
}

type grpcService struct {
	service micro.Service
}

func (s *grpcService) Register(f func(service micro.Service) error) {
	util.RegMust(func() error {
		return f(s.service)
	})
}

func (s *grpcService) Run() {
	util.PanicWhenError(s.service.Run(), "服务启动失败")
}
