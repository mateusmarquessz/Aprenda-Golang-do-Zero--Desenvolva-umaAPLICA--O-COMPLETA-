package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	enderecoExemplo := endereco{
		logradouro: "Exemplo1",
		numero:     1,
	}
	//Criando Usuario
	var u usuario
	u.nome = "Mateus"
	u.idade = 21
	fmt.Println(u)

	//Outra maneira de criar usuario
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)
}
