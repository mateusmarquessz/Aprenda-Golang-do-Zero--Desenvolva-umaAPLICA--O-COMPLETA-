package main

import "fmt"

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//Criando variavel com condicao
	//Infinit
	if outNumero := numero; outNumero > 0 {
		fmt.Println("Numero maior que Zero")
	} else if numero < -10 {
		fmt.Println("Numero menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	//Nao se pode acessa a variavel pois ela esta dentro do escopo do if
}
