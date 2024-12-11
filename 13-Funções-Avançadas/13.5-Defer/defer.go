package main

import "fmt"

func f1() {
	fmt.Println("Func 1")
}
func f2() {
	fmt.Println("Func 2")
}

func alunoAprroved(n1, n2 int) bool {
	defer fmt.Println("Média calculada")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false

}

func main() {

	defer f1()
	f2()
	aprroved := alunoAprroved(5, 8)
	fmt.Println(aprroved)

}
