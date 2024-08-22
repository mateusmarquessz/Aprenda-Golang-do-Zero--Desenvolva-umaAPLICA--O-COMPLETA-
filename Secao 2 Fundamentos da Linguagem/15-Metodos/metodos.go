package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// Criando Metodo
func (u usuario) salvar() {
	fmt.Println("dentro do metodo salvar")
}

func main() {

	usuario1 := usuario{"Usuario 1", 20}
	usuario1.salvar()

}
