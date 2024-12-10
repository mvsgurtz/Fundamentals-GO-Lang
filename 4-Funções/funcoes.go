package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	//f("func")
	result := f("func")
	fmt.Println(result)

	result1, _ := calculos(10, 20)

	fmt.Println(result1)

}
