package main

import "fmt"

func clsr() func() {
	text := "dentro do closure"

	function := func() {
		fmt.Println(text)
	}

	return function

}

func main() {
	text := "dentro do main"
	fmt.Println(text)

	func1 := clsr()
	func1()

}
