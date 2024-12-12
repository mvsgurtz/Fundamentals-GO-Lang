package enderecos

import "testing"

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTest := []cenarioDeTest{
		{"Rua Parapua", "RUA"},
		{"Avenida Parapua", "AVENIDA"},
		{"Parapua", "Invalid"},
		{"RUA do meli", "RUA"},
		{"AVENIDA do meli", "AVENIDA"},
		{" ", "Invalid"},
	}

	for _, cenario := range cenariosDeTest {
		retorno := TipoDeEndereco(cenario.enderecoInserido)

		if retorno != cenario.retornoEsperado {
			t.Errorf("Tipo %s Inválido", cenario.enderecoInserido)
		}

	}

}
