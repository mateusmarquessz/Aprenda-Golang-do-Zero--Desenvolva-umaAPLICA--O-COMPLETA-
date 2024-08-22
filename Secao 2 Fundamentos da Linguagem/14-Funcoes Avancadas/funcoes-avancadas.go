package main

import "fmt"

//Retorno Nomeado
//Aqui voce ja cria as variaveis de inicio
func calculosMatematicos(n1, n2 int) (somas int, subtracao int) {
	somas = n1 + n2
	subtracao = n1 + n2
	return
}

//Funcoes Variaticas
//ela vai receber de 1 a n int numeros
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	somas, subtracao := calculosMatematicos(10, 20)
	fmt.Println(somas, subtracao)

	totalSoma := soma(1, 2, 3, 0)
	fmt.Println(totalSoma)
}
