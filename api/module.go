package api

type Module interface {

	Init() error

	Dispatch(s string) string
}

