package main

import (
	"fmt"
)

func main() {
	//i := 0
	//for i < 10 {
	//	i++
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//}
	//
	//fmt.Println(i)
	//
	//for j := 0; j < 10; j++ {
	//	fmt.Println("Incrementando i")
	//	time.Sleep(time.Second)
	//	fmt.Println(j)
	//}

	nomes := [3]string{"Marcos", "Alves", "Goulart"}

	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for i, letra := range "TESTE" {
		fmt.Println(i, string(letra))
	}

	usr := map[string]string{
		"nome":      "Marcos",
		"sobrenome": "Goulart",
	}

	for k, value := range usr {
		fmt.Println(k, value)
	}

}
