package main

import "fmt"

type address struct {
	publicPlace string
	number      int16
}

type user struct {
	name    string
	age     int8
	address address
}

func main() {
	var user1 user
	fmt.Println("Using Structs")

	user1.name = "Luiz"
	user1.age = 20

	//defaultAdress := address{"Rua Guararapes", 86}

	user2 := user{"Luiz", 20, address{"Rua Guararapes", 86}}
	user3 := user{age: 20}

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
}
