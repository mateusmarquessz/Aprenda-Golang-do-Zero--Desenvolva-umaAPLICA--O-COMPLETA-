package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Rex", "raca": "viralata", "idade": 16}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
