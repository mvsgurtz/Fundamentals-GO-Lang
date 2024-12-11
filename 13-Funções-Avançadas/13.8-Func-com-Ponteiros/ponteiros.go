package main

import "fmt"

// Parametro por copia
func inverterSinal(num int) int {
	return num * -1
}

// Parametro por referência
func inverterSinalComPonteiro(num *int) {
	*num = *num * -1
}

func main() {

	num := 20

	numChanged := inverterSinal(num)

	fmt.Println(numChanged)
	fmt.Println(num)

	num2 := 10
	fmt.Println(num2)
	inverterSinalComPonteiro(&num2)
	fmt.Println(num2)

}
