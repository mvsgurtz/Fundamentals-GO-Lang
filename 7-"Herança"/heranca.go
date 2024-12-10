package main

import "fmt"

type pessoa struct {
	name   string
	idade  int8
	gender string
}

type estudante struct {
	pessoa
	subject string
}

func main() {

	p1 := pessoa{
		"Marcos",
		21,
		"male",
	}

	e1 := estudante{
		p1,
		"CC",
	}

	fmt.Println(e1.name)

}
