package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/aeramu/example/mocks"
	"github.com/aeramu/example/service/api"
	"testing"
)

var (
	adapter Adapter
	mockFooRepository *mocks.FooRepository
	mockBarClient *mocks.BarClient
)

func initTest()  {
	mockFooRepository = new(mocks.FooRepository)
	mockBarClient = new(mocks.BarClient)
	adapter = Adapter {
		FooRepository: mockFooRepository,
		BarClient: mockBarClient,
	}
}

func Test_service_Foo(t *testing.T)  {
	var (
		ctx = context.Background()
	)
	type args struct {
		ctx context.Context
		req api.FooReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.FooRes
		wantErr bool
	}{
		{
			name:    "should error",
			prepare: func(){

			},
			args:    args{
				ctx: ctx,
				req: api.FooReq{},
			},
			want:    nil,
			wantErr: true,
		},
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
			got, err := s.Foo(tt.args.ctx, tt.args.req)
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

func Test_service_Bar(t *testing.T)  {
	var (
		ctx = context.Background()
	)
	type args struct {
		ctx context.Context
		req api.BarReq
	}
	tests := []struct {
		name    string
		prepare func()
		args    args
		want    *api.BarRes
		wantErr bool
	}{
		{
			name:    "should error",
			prepare: func(){

			},
			args:    args{
				ctx: ctx,
				req: api.BarReq{},
			},
			want:    nil,
			wantErr: true,
		},
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
			got, err := s.Bar(tt.args.ctx, tt.args.req)
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
