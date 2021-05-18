package main

import "fmt"

type employee struct {
	salary float32
}

func main() {
	employee := employee{3000}
	salary := employee.giveRaise(0.25)
	fmt.Printf("The salary now is %f\n", salary)
}

func (e employee) giveRaise(ratio float32) float32 {
	return e.salary * (1 + ratio)
}