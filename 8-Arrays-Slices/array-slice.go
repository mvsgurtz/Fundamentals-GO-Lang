package main

import "fmt"

func main() {
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	arr2 := [3]string{"p1", "p2", "p3"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 5}
	fmt.Println(arr3)

	slc1 := []int{1, 3, 4}
	fmt.Println(slc1)

	slc1 = append(slc1, 5)
	fmt.Println(slc1)

	slc2 := arr2[1:3]
	fmt.Println(slc2)
}
