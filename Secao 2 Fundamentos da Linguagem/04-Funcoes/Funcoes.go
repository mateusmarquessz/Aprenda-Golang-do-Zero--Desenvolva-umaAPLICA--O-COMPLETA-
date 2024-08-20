package main

import "fmt"

//Funcao somar com 2 parametros, + return, por isso o terceiro int
func somar(n1 int, n2 int) int {
	return n1 + n2
}

//As funcoes podem ter mais de um retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	substracao := n1 - n2
	return soma, substracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//Declarando funcao dentro de outra funcao
	var f = func(txt string) string {
		fmt.Println("Funcao F")
		return txt
	}

	resultado := f("Texto da Funcao 1")
	fmt.Println(resultado)

	//2 variaveis recebendo 2 valores para a funcao com 2 return
	//_ para usar a funcao que somente usando uma variavel
	resultadosSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosSoma)
}
