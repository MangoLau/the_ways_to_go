package main

import "fmt"

type Request struct {
	a, b int
	replyc chan int
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func (r *Request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}

	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer(func(a, b int) int {
		return a+b
	})
	req1 := &Request{3, 4, make(chan int)}
	req2 := &Request{150, 250, make(chan int)}
	adder <- req1
	adder <- req2
	fmt.Println(req1)
	fmt.Println(req2)
	quit <- true
	fmt.Println("done")
}