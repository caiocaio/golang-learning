package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplcacao := 10 * 5
	restoDaDvidisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplcacao, restoDaDvidisao)

	//usar mesmo tipo para fazer qlq operação
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//operadores de atribuição:
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais:
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//operadores logicos
	fmt.Println("----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//operadores unarios

	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 15 //numero = numero + 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 20 //numero  numero -20
	fmt.Println(numero)
	numero *= 3 //numero = numero *3
	fmt.Println(numero)
	numero /= 10 //numero = dividido por 3
	fmt.Println(numero)
	numero %= 3 //numero = resto da divisao por 3

}
