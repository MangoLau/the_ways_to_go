package main

import (
	"fmt"
	"github.com/mangolau/the_ways_to_go/even"
)

func main() {
	var num int64 = 12
	isOdd := even.IsOdd(num)
	if isOdd {
		fmt.Printf("%d is odd.", num)
	} else {
		fmt.Printf("%d is enen.", num)
	}
}