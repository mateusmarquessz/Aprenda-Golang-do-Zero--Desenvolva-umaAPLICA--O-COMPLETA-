package main

import (
	"errors"
	"fmt"
)

func main() {
	//Possui 4 tipos de inteiros
	//A diferenca entre eles e o tamanho de bits que suporta o inteiro
	//int8, int16, int32, int64
	//possui tbm o int sozinho, e ele usuario a arquitetura do seu computador.
	//maneira explicida

	//Numeros Inteiros
	var numero int16 = 100
	//maneira implicida
	numero1 := 100

	fmt.Println(numero)
	fmt.Println(numero1)

	//int sem sinal
	// unsygned int
	var numero3 uint32 = 10000
	fmt.Println(numero3)

	//alias
	//int32 === rune
	var numero4 rune = 123456
	fmt.Println(numero4)

	//uint8 == byte
	var numero5 byte = 123
	fmt.Println(numero5)

	//Numeros Reais
	//Numeros Reais so possuiem 2 tipos
	//float32, float64
	//Segue a mesma logica de 32 bits e 64 bits
	//nao se pode usar float sem os 32 ou 64
	var numeroReal float32 = 123.43
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12222222223.43
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.65
	fmt.Println(numeroReal3)

	//Strings
	var str string = "Texto"
	fmt.Println(str)

	srt2 := "Texto2"
	fmt.Println(srt2)

	//Ele vai pegar o numero da tabela Hash
	char := 'B'
	fmt.Println(char)

	//Valor 0
	//Valor que e atribuido na variavel caso voce nao coloque valor

	//Booleano
	//true ou false
	//O valor 0 do booleano e false
	var Booleano1 bool = true
	fmt.Sprintln(Booleano1)

	//Tipo Error
	//<nil>
	//Declarando erro
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
