package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:81") // tcp ipv4
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "https://www.all4learn.com") // udp
	checkConnection(conn, err)
	//conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // tcp ipv6
	//checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
}