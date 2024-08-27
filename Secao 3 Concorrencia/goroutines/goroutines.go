package main

import (
	"fmt"
	"time"
)

func main() {
	//Cocorrencia != paralelismo

	//Go serve para seguir o programa normalmente
	go escrever("Ola Mundo!")
	escrever("Programando em Go!")
}

func escrever(string) {
	for {
		fmt.Println("Ola Mundo")
		time.Sleep(time.Second)
	}
}
