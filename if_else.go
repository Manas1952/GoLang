package main

import "fmt"

func main() {
	if num := 4; num%3 == 0 {
		fmt.Println("divisible by 2")
	} else if num > 0 {
		fmt.Println("greater than 0")
	} else {
		fmt.Println("in else")
	}
}
