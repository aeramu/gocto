package service

type Adapter struct {
	FooRepository FooRepository
	BarClient BarClient
}

type FooRepository interface {
}

type BarClient interface {
}
