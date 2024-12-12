package main

import (
	"fmt"
	"introducao/enderecos"
)

func main() {
	tipoDeEndereco := enderecos.TipoDeEndereco("Rua Parapua")

	fmt.Println(tipoDeEndereco)
}
