package main

import (
	"errors"
	"fmt"
)

func main() {
	// Números inteiros
	var number1 int8 = 100
	var number2 int16 = 10000
	var number3 int32 = 1000000000
	var number4 int64 = 1000000000000000000
	var number7 uint8 = 100 // Não pode ser negativo

	// Números Reais
	var number5 float32 = 100.50
	var number6 float64 = 100.50

	// String
	var texto string = "Isto é uma string"
	texto2 := "Isto também é uma string"
	texto3 := 'B' // Isto não é um char, não existe Char em Go

	// Booleanos
	var boolean1 bool = true
	var boolean2 bool = false

	// Erro
	var err error = errors.New("Isto é um erro")

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(number6)
	fmt.Println(number7)
	fmt.Println(texto)
	fmt.Println(texto2)
	fmt.Println(texto3)
	fmt.Println(boolean1)
	fmt.Println(boolean2)
	fmt.Println(err)

}
