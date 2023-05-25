package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println(array)

	array[4] = 5
	fmt.Println([3]string{"1", "3"})

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("array 'b':", b, ", length:", len(b))
	fmt.Println()

	var twoD [2][3]int

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
