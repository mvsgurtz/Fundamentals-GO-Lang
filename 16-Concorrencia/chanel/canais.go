package main

import (
	"fmt"
	"time"
)

func main() {
	chanel := make(chan string)
	fmt.Println("Inciando Progama.")
	go escrever("Funcionando", chanel)

	for mensagem := range chanel {
		fmt.Println(mensagem)
	}

	fmt.Println("Programa finalizado.")

}

func escrever(txt string, chanel chan string) {
	for i := 0; i < 5; i++ {
		chanel <- txt
		time.Sleep(time.Second)
	}

	close(chanel)
}
