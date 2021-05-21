package main

import "fmt"

type Empty interface {}
type semaphore chan Empty

func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<- s
	}
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c)
}

func sum(x, y int, c chan int) {
	c <- x + y
}