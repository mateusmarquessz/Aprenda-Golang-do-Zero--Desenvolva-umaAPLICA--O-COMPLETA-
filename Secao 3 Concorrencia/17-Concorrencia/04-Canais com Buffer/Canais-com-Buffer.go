package main

import "fmt"

func main() {
	//A diferenca do canal com Buffer e o canal normal, e que ele so vai encher quando bater o limite dele, por isso nao da erro de codigo
	canal := make(chan string, 2)
	canal <- "Ola mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
