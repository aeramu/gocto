package api

type FooReq struct {
}

type FooRes struct {
}

type BarReq struct {
}

type BarRes struct {
}

func (req FooReq) Validate() error {
	return nil
}

func (req BarReq) Validate() error {
	return nil
}
