// Teste de Unidade
package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Sintaxe de um teste
func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Praca das Rosas", "Tipo Invalido"},
		{"Entrada Qualquer", "Estrada"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s e diferente do esperado %s", tipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}
}
