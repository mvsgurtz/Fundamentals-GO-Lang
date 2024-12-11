package main

import "fmt"

func calc(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

func escrever(text string, nums ...int) {
	for _, num := range nums {
		fmt.Println(text, num)
	}
}

func main() {
	calc(1, 2, 3, 4)
	escrever("Número:", 1, 2, 3, 4, 5)
}
