package auxiliar

import "fmt"

//Escrever registra uma mensagem na tela
func Escrever() {
	fmt.Sprintln("Escrevendo do pacote auxiliar")
	//chamando funcao do auxiliar 2 que esta no mesmo pacote
	//referenciando pacote
	escrever2()
}
