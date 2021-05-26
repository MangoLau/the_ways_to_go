package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	p := new(Person)
	p.Name = "Lau"
	p.Age = 18
	fmt.Println(p)

	p2 := &Person{"Chris", 20}
	fmt.Println(p2)

	p3 := NewPerson("Rose", 19)
	fmt.Println(p3)
}

// 构建函数
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}