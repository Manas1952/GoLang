package main

import "fmt"

func main() {
	closure1 := myFunc()

	fmt.Println(closure1())
	fmt.Println(closure1())
	fmt.Println(closure1())

	fmt.Println()

	closure2 := myFunc()
	fmt.Println(closure2())
	fmt.Println(closure2())
	fmt.Println(closure2())
}

func myFunc() func() (int, int) {

	i := 0
	return func() (int, int) {
		i++
		return i, i + 1
	}
}
