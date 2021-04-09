package main

import "fmt"

func main() {
	var a, b = 10, 5
	sum,product, diff := cal(a, b)
	fmt.Printf("cal() a:%d, b:%d; sum:%d, product:%d, diff:%d\n", a, b, sum, product, diff)

	sum2, product2, diff2 := cal2(a, b)
	fmt.Printf("cal2() a:%d, b:%d; sum:%d, product:%d, diff:%d\n", a, b, sum2, product2, diff2)
}

func cal(a,b int) (sum, product, difference int) {
	sum, product, difference = a + b, a * b, a - b
	return
}

func cal2(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}