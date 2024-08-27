//Generator basicamente vai ser uma funcao que vai encapsular uma go rotine, e devolver um canal para gente

package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola mundo!")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// Com generator eu posso esconder a complexidade na main
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
