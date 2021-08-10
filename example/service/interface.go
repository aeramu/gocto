package service

type Adapter struct {
	BookRepository BookRepository
}

type BookRepository interface {
}
