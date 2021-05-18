package main

import (
	"fmt"
	"github.com/mangolau/the_ways_to_go/even"
)

func main() {
	for i:=0; i<= 100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}