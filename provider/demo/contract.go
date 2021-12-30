package demo

const Key = "nwfw:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}