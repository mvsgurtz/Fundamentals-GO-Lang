package enderecos

import "strings"

// TipoDeEndereco o tipo de endereço
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida"}

	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	valid := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			valid = true
		}
	}

	if valid {
		return strings.ToTitle(primeiraPalavra)
	}

	return "Invalid"

}
