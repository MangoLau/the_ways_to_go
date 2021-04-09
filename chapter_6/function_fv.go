package main

import "fmt"

func main() {
	fv := func(s string) {
		fmt.Print("Hello World")
	}
	fmt.Printf("The type is %T, the value is %v", fv, fv)
}