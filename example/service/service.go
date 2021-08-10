package service

import (
	"context"
	"github.com/aeramu/gocto/example/service/api"
)

type Service interface {
	GetBookByISBN(ctx context.Context, req api.GetBookByISBNReq) (*api.GetBookByISBNRes, error)
	GetBookByID(ctx context.Context, req api.GetBookByIDReq) (*api.GetBookByIDRes, error)
	InsertBook(ctx context.Context, req api.InsertBookReq) (*api.InsertBookRes, error)
}

func NewService(adapter Adapter) Service {
	return &service {
		adapter: adapter,
	}
}

type service struct {
	adapter Adapter
}

func (s *service) GetBookByISBN(ctx context.Context, req api.GetBookByISBNReq) (*api.GetBookByISBNRes, error) {
	panic("implement me")
}

func (s *service) GetBookByID(ctx context.Context, req api.GetBookByIDReq) (*api.GetBookByIDRes, error) {
	panic("implement me")
}

func (s *service) InsertBook(ctx context.Context, req api.InsertBookReq) (*api.InsertBookRes, error) {
	panic("implement me")
}
