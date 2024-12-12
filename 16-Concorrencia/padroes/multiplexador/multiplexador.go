package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := escrever("Hi")
	chan2 := escrever("Ola")

	chars := multiplexar(chan1, chan2)

	for char := range chars {
		fmt.Println(char)
	}

}

func multiplexar(chan1, chan2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-chan1:
				canalDeSaida <- mensagem
			case mensagem := <-chan2:
				canalDeSaida <- mensagem
			}

		}
	}()
	return canalDeSaida
}

func escrever(txt string) <-chan string {
	chanel := make(chan string)

	go func() {
		for {
			chanel <- txt
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return chanel
}
