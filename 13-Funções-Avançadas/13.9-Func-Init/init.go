package main

import "fmt"

// Pode ter 1 por arquivo, não por package
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
