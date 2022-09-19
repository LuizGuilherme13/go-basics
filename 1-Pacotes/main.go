package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello Go!")
	auxiliar.Write()
	err := checkmail.ValidateFormat("luiz@hotmail.com")

	fmt.Println(err)
}
