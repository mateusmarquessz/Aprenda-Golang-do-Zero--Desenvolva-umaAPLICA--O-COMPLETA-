package main

import (
	"fmt"
	"time"
)

func main() {
	//Canal e um canal de comunicao que serve para sincronizar nossas goroutines
	canal := make(chan string)
	go escrever("Ola Mundo!", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do Programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//canal <- texto mandando valor
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
