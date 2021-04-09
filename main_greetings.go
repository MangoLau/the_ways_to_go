package main

import (
	"./pack1"
	"fmt"
	"github.com/mangolau/the_ways_to_go/greetings"
)

func main() {
	isEvening := greetings.IsEvening()
	if isEvening {
		greetings.GoodNight()
	} else {
		greetings.GoodDay()
	}

	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
}