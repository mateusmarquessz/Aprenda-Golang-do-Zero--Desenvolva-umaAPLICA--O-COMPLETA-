package main

import (
	"log"
	"net/http"
)

func main() {
	//HTTP E UM PROTOCOLO DE COMUNICACAO - BASE DA COMUNICACAO WEB

	//CLIENTE(FAZ A REQUISICAO) - SERVIDOR (PROCESSA REQUISICAO E ENVIA A RESPOSTA)

	//Request - Response

	//Rotas
	//URI - Identificador do Recurso
	//Metodo - GET , POST, PUT, DELETE
	//Get - Busca dados
	//Post - Cadastrar dados
	//Put - Atualizar dados
	//DELETE - deleta dados

	//Criando um servidor para http

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola Mundo!"))
	})

	http.HandleFunc("/Usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Lista de Usuarios"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
