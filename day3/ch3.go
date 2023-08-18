package main

type handle1 interface {
	handle()
}
type handle2 interface {
	handle()
}
type person struct {
	name string
}

func (p person) handle() {

}
