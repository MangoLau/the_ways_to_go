package main

import (
	"fmt"
)

func main() {
	s := Simple{33}
	fmt.Printf("The value is: %d\n", s.Get());
	fI(s)
	rs := RSimple{44}
	gI(rs)
	gI(22.0)

}

type Simpler interface {
	Get() int32
	Set(v int32)
}

type Simple struct {
	v int32
}

func (s *Simple) Set(v int32) {
	s.v = v
}

func (s *Simple) Get() int32 {
	return s.v
}

type RSimple struct {
	v int32
}

func (s *RSimple) Set(v int32) {
	s.v = v
}

func (s *RSimple) Get() int32 {
	return s.v
}

func fI(s interface{}) {
	switch s.(type) {
		case Simple:
			fmt.Printf("Param %v is Simple\n", s)
		case RSimple:
			fmt.Printf("Param %v is RSimple\n", s)
		default:
			fmt.Printf("Unknown type\n")
	}
}

func gI(any interface{}) {
	switch any.(type) {
		case Simple:
			fmt.Printf("Param %v is Simple\n", any)
		case RSimple:
			fmt.Printf("Param %v is RSimple\n", any)
		case int:
			fmt.Printf("Param %v is integer\n", any)
		default:
			fmt.Printf("Unknown type\n")
	}
}