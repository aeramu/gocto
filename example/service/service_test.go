package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/aeramu/gocto/example/mocks"
	"github.com/aeramu/gocto/example/service/api"
	"testing"
)

var (
	adapter Adapter
	mockBookRepository *mocks.BookRepository
)

func initTest()  {
	mockBookRepository = new(mocks.BookRepository)
	adapter = Adapter {
		BookRepository: mockBookRepository,
	}
}

func Test_service_GetBookByISBN(t *testing.T)  {
	var (

	)
	type args struct {
		ctx context.Context
		req api.GetBookByISBNReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.GetBookByISBNRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.GetBookByISBN(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_service_GetBookByID(t *testing.T)  {
	var (

	)
	type args struct {
		ctx context.Context
		req api.GetBookByIDReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.GetBookByIDRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.GetBookByID(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_service_InsertBook(t *testing.T)  {
	var (

	)
	type args struct {
		ctx context.Context
		req api.InsertBookReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.InsertBookRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initTest()
			if tt.prepare != nil {
				tt.prepare()
			}
			s := &service{
				adapter: adapter,
			}
			got, err := s.InsertBook(tt.args.ctx, tt.args.req)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
