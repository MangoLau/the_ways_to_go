package main

import "fmt"

type Address struct {
	address string
}

type VCard struct {
	name string
	birthday string
	avatar string
	address *Address
}

func main() {
	address := Address{"address"}
	card := VCard{"Name","2021-04-12","url",&address}
	fmt.Printf("The name is %s\n", card.name)
	fmt.Printf("The birthday is %s\n", card.birthday)
	fmt.Printf("The avatar is %v\n", card.avatar)
	fmt.Printf("The address is %s\n", card.address.address)
	fmt.Println(card)
}