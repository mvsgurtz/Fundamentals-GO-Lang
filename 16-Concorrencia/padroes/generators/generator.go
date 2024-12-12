package main

import (
	"fmt"
	"time"
)

func main() {

	canal := escrever("teste1")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

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
