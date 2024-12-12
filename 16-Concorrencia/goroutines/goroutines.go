package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Escrever 1")
	escrever("Escrever 2")

}

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}

}
