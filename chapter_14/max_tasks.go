package main

import "fmt"

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b int
	replyc chan int
}

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

func process(r *Request) {
	// do something
	r.replyc <- r.a + r.b
}

func handle(r *Request) {
	sem <- 1 // doesn't matter what we put in it
	process(r)
	<-sem // one empty place in the buffer: the next request can start
}

func server(service chan *Request) {
	for {
		request := <-service
		go handle(request)
	}
}

func main() {
	service := make(chan *Request)
	go server(service)
	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}
	service <- req1
	service <- req2
	fmt.Println(req1)
	fmt.Println(req2)
	fmt.Println("done")
}