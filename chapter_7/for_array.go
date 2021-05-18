package main

import "fmt"

func main() {
	var arr [15]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}
}