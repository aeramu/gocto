package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/aeramu/example/service/api"
	"testing"
)

func Test_service_Foo(t *testing.T)  {
	type fields struct {
		adapter Adapter
	}
	type args struct {
		ctx context.Context
		req api.FooReq
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func()
		args    args
		want    *api.FooRes
		wantErr bool
	}{
		{
			name:    "should success",
			fields:  fields{
				adapter: Adapter{

				},
			},
			prepare: func(){
				
			},
			args:    args{
				ctx: context.Background(),
				req: api.FooReq{},
			},
			want:    &api.FooRes{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				adapter: tt.fields.adapter,
			}
			tt.prepare()
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
	type fields struct {
		adapter Adapter
	}
	type args struct {
		ctx context.Context
		req api.BarReq
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func()
		args    args
		want    *api.BarRes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				adapter: tt.fields.adapter,
			}
			tt.prepare()
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
