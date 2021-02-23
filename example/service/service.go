package service

import (
	"context"
	"github.com/aeramu/example/service/api"
)

type Service interface {
	Foo(ctx context.Context, req api.FooReq) (*api.FooRes, error)
	Bar(ctx context.Context, req api.BarReq) (*api.BarRes, error)
}

func NewService(adapter Adapter) Service {
	return &service {
		adapter: adapter,
	}
}

type service struct {
	adapter Adapter
}

func (s *service) Foo(ctx context.Context, req api.FooReq) (*api.FooRes, error) {
	panic("implement me")
}

func (s *service) Bar(ctx context.Context, req api.BarReq) (*api.BarRes, error) {
	panic("implement me")
}
