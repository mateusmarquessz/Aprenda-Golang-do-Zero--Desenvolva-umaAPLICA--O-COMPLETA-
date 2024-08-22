package main

import "fmt"

//Ponteiros funcionam de forma que armazenam memoria
func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	//aqui o valor da variavel2 nao muda
	fmt.Println(variavel1, variavel2)

	//Ponteiro e uma referencia de memoria
	var variavel3 int
	var ponteiro *int //armazena o local da memoria para um inteiro

	variavel3 = 100
	ponteiro = &variavel3 //recebendo o valor da variavel3 armazenando nesse espaco de memoria

	//chamando o ponteiro com o endereco de memoria com o valor da variavel3
	fmt.Println(variavel3, *ponteiro)

	variavel3 = 150
	//aqui muda pois usamos o endereco da memoria
	fmt.Println(variavel3, *ponteiro)
}
