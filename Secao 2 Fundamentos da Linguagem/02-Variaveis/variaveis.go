package main

import "fmt"

//A duas maneiras de declarar variaveis
//De forma implícita ou explicita
//Go e fortemente tipado
//Go nao deixa usar variaveis que voce nao vai usar
func main() {
	//Forma explicita
	var variavel1 string = "Variavel 1"
	//Forma implicita
	variavel2 := "Variavel2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Forma explicita
	var (
		variavel3 string = "a"
		variavel4 string = "a"
	)

	fmt.Println(variavel3, variavel4)

	//Forma implicita
	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante1"
	fmt.Println(constante1)

	//Mudando os valores das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
