package main

import "fmt"

func main() {
	//Aritmeticos e o padrao de qualquer linguagem
	//Go nao deixa fazer a conta com tipos diferentes de int
	var n1 int16 = 10
	var n2 int32 = 25
	soma := n1 + int16(n2)
	fmt.Println(soma)

	//Atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//Operadores Relacionais
	//Sempre retorna true ou false
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	//Diferente
	fmt.Println(1 != 2)

	//Operadores Logicos
	//AND
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	//OU
	fmt.Println(verdadeiro || falso)
	//NEGACAO
	fmt.Println(!verdadeiro)

	//Operadores Unarios
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 1 //numero = numero - 1

	numero *= 3 //numero = numero * 3

	fmt.Println(numero)

	//Operador Ternario

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
