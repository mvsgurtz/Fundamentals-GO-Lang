package main

import "fmt"

func main() {

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)
	v1++
	fmt.Println(v1, v2)

	//	Ponteiro:
	var v3 int
	var pont1 *int

	v3 = 100
	pont1 = &v3

	fmt.Println(v3, pont1)
	//Desreferenciação
	fmt.Println(*pont1)

	v3 = 150
	fmt.Println(v3, pont1)
	fmt.Println(*pont1)

	//	Arrays Internos
	fmt.Println("----------")
	slc2 := make([]float32, 10, 11)
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

	slc2 = append(slc2, 10)
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))
	
	slc2 = append(slc2, 12)
	fmt.Println(slc2)
	fmt.Println(len(slc2))
	fmt.Println(cap(slc2))

}
