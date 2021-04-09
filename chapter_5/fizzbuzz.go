package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(num int) {
	remainder3 := num % 3
	remainder5 := num % 5
	switch {
		case remainder3 == 0 && remainder5 == 0:
			fmt.Print("FizzBuzz\n")
		case remainder3 == 0:
			fmt.Print("Fizz\n")
		case remainder5 == 0:
			fmt.Print("Buzz\n")
		default:
			fmt.Printf("%d\n", num)
	}
}
