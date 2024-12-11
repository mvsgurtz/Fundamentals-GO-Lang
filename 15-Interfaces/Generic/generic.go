package main

import "fmt"

func generica(intf interface{}) {
	fmt.Println(intf)
}

func main() {
	generica("String")
	generica(12)
	generica(false)

}
