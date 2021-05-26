package main

import (
	"errors"
)

func main() {
	// 如何在程序出错时终止程序：
	err := makeError()
	/*if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}*/
	if err != nil {
		panic("ERROR occurred: " + err.Error())
	}
}

func makeError() error {
	return errors.New("error maked")
}