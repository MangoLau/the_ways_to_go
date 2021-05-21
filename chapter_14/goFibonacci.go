package main

import "fmt"

func main() {
	ch := make(chan int)
	go fibonterms(15, ch)
	var i = 0
	for {
		if result, ok := <- ch; ok {
			fmt.Printf("fibonacci(%d) is %d\n", i, result)
		} else {
			break
		}
		i++
	}
}

func fibonterms(n int, ch chan int) {
	for i := 0; i <= n; i++ {
		ch <- fibonacci(i)
	}
	close(ch)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 2) + fibonacci(n - 1)
	}
	return
}