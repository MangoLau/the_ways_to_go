package main

import "fmt"

func main()  {
	forFun(15)
	gotoFun(15)
}

func forFun(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("the i is %d\n", i)
	}
}

func gotoFun(count int)  {
	i := 0
	START:
		fmt.Printf("the i is %d\n", i)
	i++
	if i < count {
		goto START
	}
}
