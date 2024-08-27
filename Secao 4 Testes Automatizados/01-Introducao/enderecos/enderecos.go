package enderecos

import "strings"

//TipoDeEndereco
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua", "avenida", "estrada", "rodovia",
	}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavradoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	//Rua dos Bobos
	//["Rua", "Dos", "Bobos"]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavradoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavradoEndereco
	}

	return "Tipo Invalido"
}
