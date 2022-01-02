package demo

const DemoKey = "nwfw:demo"

type IService interface {
	GetAllStudent() []Student
}


type Student struct {
	ID int
	Name string
}