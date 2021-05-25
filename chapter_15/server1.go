package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("Starting the server ...")

	// create listener:
	listener, error = net.Listen("tcp", "localhost:50000")
	checkError(error)
	// listen and accept connections from clients:
	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff2(conn)
	}
}

func doServerStuff2(conn net.Conn) {
	var buf []byte
	var error error

	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		checkError(error)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}
		// op commando WHO: write out mapUsers
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}
		// extract clientname:
		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		//fmt.Printf("The clientname is ---%s---\n", string(clName))
		// set clientname active in mapUsers:
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

func checkError(err error)  {
	if err != nil {
		panic("Error: " + err.Error()) // terminate program
	}
}

func DisplayList()  {
	fmt.Println("------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("------------------------------------")
}