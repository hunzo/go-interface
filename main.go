package main

import "fmt"

type getdata interface {
	show() interface{}
}

type sString struct {
	s string
}

type iInteger struct {
	i int
}

func (s sString) show() interface{} {
	return s.s
}

func (i iInteger) show() interface{} {
	return i.i
}

func call(g getdata) {
	fmt.Println(g.show())
}

func main() {
	s := sString{
		s: "show String",
	}
	i := iInteger{
		i: 1000,
	}

	call(s)
	call(i)

}
