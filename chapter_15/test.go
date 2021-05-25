package main

import (
	"fmt"
	"net/url"
)

func main()  {
	u, err := url.ParseRequestURI("https://www.baidu.com")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(u)
}