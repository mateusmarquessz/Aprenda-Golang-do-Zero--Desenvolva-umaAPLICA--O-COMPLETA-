package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	//puxando a funcao de escrever, pelo pacote auxiliar
	//nao se deve colocar letra minuscula, da erro de sintaxe
	//so se pode letra minuscula em arquivos de mesmo pacote
	auxiliar.Escrever()

	//usando de um pacote externa
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
