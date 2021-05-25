package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	var counter = 100000
	//str := "a";
	//start := time.Now()
	//for i := 0; i < counter; i++ {
	//	str += " ";
	//}
	//fmt.Println(str)
	//end := time.Now()
	//elapsed := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", elapsed)

	str2 := ""
	var b bytes.Buffer
	b.WriteString("b")
	start2 := time.Now()
	for i := 0; i < counter; i++ {
		b.WriteString("")
	}
	str2 = b.String()
	end2 := time.Now()
	fmt.Println(str2)
	elapsed2 := end2.Sub(start2)
	fmt.Printf("longCalculation took this amount of time: %s\n", elapsed2)
}