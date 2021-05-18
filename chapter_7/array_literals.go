package main

import "fmt"
func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	//var arrKeyValue = [5]string{3: "chris", 4: "Ron"}
	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrAge[i])
	}
}