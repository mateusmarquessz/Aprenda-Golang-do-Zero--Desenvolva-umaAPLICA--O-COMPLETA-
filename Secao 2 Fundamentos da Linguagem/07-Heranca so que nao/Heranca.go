package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// heranca entre muitas aspas
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	
	p1 := pessoa{"Joao", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Letras", "Uemg"}
	fmt.Println(e1)
}
