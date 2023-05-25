package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for index, value := range nums {
		fmt.Println(index, value)
		sum += value
	}

	fmt.Println(sum)
	fmt.Println()

	m := map[int]string{1: "apple", 2: "banana"}
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println()

	for index, char := range "Manas" {
		fmt.Println(index, char)
	}
}
