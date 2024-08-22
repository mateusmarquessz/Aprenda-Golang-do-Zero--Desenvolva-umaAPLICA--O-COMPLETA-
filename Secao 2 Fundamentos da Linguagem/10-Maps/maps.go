package main

import "fmt"

func main() {
	fmt.Println("Maps")

	//Dentro dos [] e o tipo das chaves fora dos [] e tipo dos valores
	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
		"curso": {
			"campus": "1",
			"curso":  "Engenharia",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
}
