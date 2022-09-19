package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	var soma int

	soma = somar(100, 100)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	f("Hello")

	resSoma, _ := calculos(10, 40)
	fmt.Println(resSoma)

}
