package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	v, err := MySqrt(10.0)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("sqrt: %f\n", v)
	fmt.Println(MySqrt2(5))
}

func MySqrt(f float64) (ret float64, err error) {
	if f < 0 {
		ret, err = float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		ret, err = math.Sqrt(f), nil
	}
	return
}

func MySqrt2(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}