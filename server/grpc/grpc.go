package grpc

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/winkb/ddup-go-util/util"
)

func Handle(service micro.Service, reg registry.Registry) *grpcService {
	if reg != nil {
		service.Init(micro.Registry(reg))
	}

	return &grpcService{
		registry: reg,
		service:  service,
	}
}

type grpcService struct {
	registry registry.Registry
	service  micro.Service
}

func (s *grpcService) GetService() micro.Service {
	return s.service
}

func (s *grpcService) Register(f func(service micro.Service) error) {
	util.RegMust(func() error {
		return f(s.service)
	})
}

func (s *grpcService) Run() {
	util.PanicWhenError(s.service.Run(), "服务启动失败")
}
