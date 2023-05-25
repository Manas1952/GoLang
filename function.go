package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(add(1, 2, "Manas"))

	a, b, c := vals()
	fmt.Println(a, b, c)

	fmt.Println(sums(1, 2, 3, 4, 5))
}

func sum(a int, b int) int {
	return a + b
}

func add(a, b int, s string) [3]int {
	c := [3]int{1, 2, 3}
	return c
}

func vals() (int, string, [1]bool) {
	return 1, "two", [1]bool{true}
}

func sums(nums ...int) int {

	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
