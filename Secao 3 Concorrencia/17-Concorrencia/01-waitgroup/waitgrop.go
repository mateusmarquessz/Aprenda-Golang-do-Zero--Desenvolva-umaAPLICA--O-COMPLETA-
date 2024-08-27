package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Cocorrencia != paralelismo

	//E uma maneira de sincronizar as funcoes goroutines
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo!")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //2
}

func escrever(string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Ola Mundo")
		time.Sleep(time.Second)
	}
}
