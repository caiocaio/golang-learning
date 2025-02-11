package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	//assim só uma variavel vai mudar de valor.
	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiro é uma referencia de memoria:
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	//assim estou alterando o valor dentro do endereço de memoria!
	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)

}
