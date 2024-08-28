package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// Passando Struct para json usando Marshal, e newBuffer para eu ver os dados
func main() {
	c := cachorro{"Rex", "Viralata", 3}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	fmt.Println(bytes.NewBuffer(cachorroEmJson))
}
