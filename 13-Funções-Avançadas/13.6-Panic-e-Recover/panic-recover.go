package main

import "fmt"

func recuperar() {
	fmt.Println("Tentativa de recuperar")
	if r := recover(); r != nil {
		fmt.Println("Recuperado")
	}
}

func approval(n1, n2 float64) bool {
	defer recuperar()

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//Para a execução do programa
	panic("Média é 6")
}

func main() {

	fmt.Println(approval(6, 6))

}
