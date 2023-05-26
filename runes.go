package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str = "สวัสดี"
	//var str1 = "Manas"

	fmt.Println("Length:", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(str))
}
