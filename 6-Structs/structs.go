package main

import "fmt"

type user struct {
	name  string
	idade int8
}

func main() {
	fmt.Println("Struct")

	u := user{
		"marcos",
		21,
	}

	u2 := user{name: "Marquitos"}
	
	fmt.Println(u2)
	fmt.Println(u)
	u.name = "jorge"
	fmt.Println(u)
}
