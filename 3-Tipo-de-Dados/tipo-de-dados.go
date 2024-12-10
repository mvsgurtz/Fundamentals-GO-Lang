package main

import (
	"errors"
	"fmt"
)

func main() {
	var num int32 = 1000
	fmt.Println(num)
	var num1 int = 200
	fmt.Println(num1)
	num2 := 100000
	fmt.Println(num2)
	var num3 uint = 200
	fmt.Println(num3)

	//	Alias

	// rune = int32
	var num4 rune = 1234
	// int8 = byte
	var num5 byte = 14
	fmt.Println(num4, num5)

	var num6 float32 = 134.4
	num7 := 1234.3
	fmt.Println(num6, num7)

	//Caracter na tabela ASCII
	char := 'A'
	fmt.Println(char)

	//	Boolean
	var boolean1 bool
	fmt.Println(boolean1)

	//	Error
	var erro error = errors.New("Error")
	fmt.Println(erro)
}
