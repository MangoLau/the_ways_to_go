package main

import (
	"fmt"
)

func main() {
	s := Simple{33}
	fmt.Printf("The value is: %d\n", s.Get());
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