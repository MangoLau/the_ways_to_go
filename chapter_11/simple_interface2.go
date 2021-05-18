package main

import (
	"fmt"
)

func main() {
	s := Simple{33}
	fmt.Printf("The value is: %d\n", s.Get());
	fi(s)

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

func fi(s interface{}) {
	switch s.(type) {
		case Simple:
			fmt.Printf("Param %v is Simple\n", s)
		case RSimple:
			fmt.Printf("Param %v is RSimple\n", s)
		default:
			fmt.Printf("Unknown type\n")
	}
}