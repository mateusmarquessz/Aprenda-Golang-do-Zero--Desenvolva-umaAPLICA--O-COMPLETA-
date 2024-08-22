package main

import (
	"fmt"
)

func main() {
	// i := 0

	// //Seria como um while
	// for i < 10 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando i", i)
	// }

	//For padrao de toda linguagem
	// for j := 0; j < 10; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando j", j)
	// }

	// nomes := [3]string{"Joao", "Mateus", "Rafael"}

	// //For para Strings
	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		//Loop infinito
	}
}
