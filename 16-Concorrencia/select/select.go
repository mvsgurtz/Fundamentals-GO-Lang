package main

import (
	"fmt"
	"time"
)

func main() {
	chanel1, chanel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			chanel1 <- "Chanel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			chanel2 <- "Chanel 2"
		}
	}()

	for {
		select {
		case mensagem1 := <-chanel1:
			fmt.Println(mensagem1)
		case mensagem2 := <-chanel2:
			fmt.Println(mensagem2)
			
		}

	}
}
