package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func (u user) salvar() {
	fmt.Printf("Salvando %s no banco de dados\n", u.nome)
}

func (u *user) fazerAniversario() {
	u.idade++
}

func main() {

	user1 := user{"Marcos", 21}

	user1.salvar()

	user1.fazerAniversario()

	fmt.Println(user1)

}
