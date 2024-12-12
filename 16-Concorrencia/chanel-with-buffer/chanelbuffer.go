package main

import "fmt"

func main() {

	chanel := make(chan string, 2)

	chanel <- "Chanel 1"
	chanel <- "Chanel 2"

	mensagem := <-chanel
	mensagem2 := <-chanel

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
