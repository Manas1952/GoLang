package main

import "fmt"

func main() {
	n := 1
	changeByValue(n)
	fmt.Println(n, &n)

	changeByReference(&n)
	fmt.Println(n, &n)

	var a int = 1
	var b *int = &a

	*b = 2
	fmt.Println(a, *b)

}

func changeByValue(n int) {
	n = 0
}

func changeByReference(ptr *int) {
	*ptr = 0
}
