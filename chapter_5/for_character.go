package main

import "fmt"

func main() {
	// 1.使用 2 层嵌套 for 循环
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
	// 2.仅用 1 层 for 循环以及字符串连接
	s := "G"
	for i := 0; i < 25; i++ {
		fmt.Printf("%s\n", s)
		s += "G"
	}
}
