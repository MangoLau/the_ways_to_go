package main

import "fmt"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

type Rectangle struct {
    length, width float32
}

func (r Rectangle) Area() float32 {
    return r.length * r.width
}

type Circle struct {
    radius float32
}

func (c Circle) Area() float32 {
    return 3.14 * c.radius * c.radius
}

func main() {
    r := Rectangle{5, 3}
    q := &Square{5}
    c := &Circle{2}

    shapes := []Shaper{r, q, c}
    fmt.Println("Looping through shapes for area ...")
    for n, _ := range shapes {
        fmt.Println("Shape details: ", shapes[n])
        fmt.Println("Area of this shape is: ", shapes[n].Area())
    }
}