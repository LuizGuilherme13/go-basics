package main

import "fmt"

func main() {
	var texto1 string = "Isto é um texto"
	texto2 := "Isto também é um texto"
	var (
		texto3 string = "Texto 3"
		texto4 string = "Texto 4"
	)
	texto5, texto6 := "Texto 5", "Texto 6"

	fmt.Println(texto1)
	fmt.Println(texto2)
	fmt.Println(texto3, texto4)
	fmt.Println(texto5, texto6)
}
