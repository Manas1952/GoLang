package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s []string
	fmt.Println(s, reflect.TypeOf(s), s == nil, len(s))

	s = make([]string, 3)
	fmt.Println(s, "length:", len(s), "capacity:", cap(s))
	fmt.Println()

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s, ", s[2]=", s[2])
	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println()
	fmt.Println(c)

	l := s[:]
	fmt.Println(l)
	fmt.Println()

	i := []string{"h", "i", "j"}
	fmt.Println("->", i)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
