package main

import "fmt"

func main() {
	print(1, 2, 4, 5, 7, 9)
}

func print(numbers ...int64)  {
	for _, v := range numbers {
		fmt.Printf("The value is %d\n", v)
	}
}