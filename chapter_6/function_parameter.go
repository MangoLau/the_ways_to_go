package main

import (
	"fmt"
	"strings"
)

func main()  {
	callback(1, Add)
	strings.Map()
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}