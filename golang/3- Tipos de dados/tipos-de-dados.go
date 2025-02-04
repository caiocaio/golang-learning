package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 10012313
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	//alias int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//alias uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float32 = 1231233.45312313
	fmt.Println(numeroReal2)

	numeroReal3 := 123123.2323
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto"
	fmt.Println(str2)

	//numero do caracter na tabela asc
	char := 'B'
	fmt.Println(char)

	//todo tipo tem um valor default, da string é vazio, do int é zero
	var texto string
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
