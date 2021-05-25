package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var in string
	fmt.Println("Please enter url: ")
	fmt.Scanln(&in)
	_, err := url.ParseRequestURI(in)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	res, err := http.Get(in)
	checkError(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Printf("GOT: %q", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}