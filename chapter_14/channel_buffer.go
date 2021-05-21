package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		time.Sleep(15 * 1e9)
		fmt.Println("received", <-ch)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}
