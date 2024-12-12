package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(2)

	go func() {
		escrever("Escrever 1")
		waitgroup.Done()
	}()
	go func() {
		escrever("Escrever 2")
		waitgroup.Done()
	}()

	waitgroup.Wait()

}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}

}
