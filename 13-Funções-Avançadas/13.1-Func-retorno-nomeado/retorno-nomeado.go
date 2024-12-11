package main

import "fmt"

func calc(n1, n2 int) (some, sub int) {
	some = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	some, sub := calc(1, 2)
	fmt.Println(some, sub)
}
