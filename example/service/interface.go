package service

import (
	"context"
)

type Adapter struct {
	FooRepository FooRepository
	BarClient BarClient
}

type FooRepository interface {
	Foo(ctx context.Context) error
}

type BarClient interface {
	Foo(ctx context.Context) error
}
