package main

import "fmt"

func main() {
	var m = make(map[string]int)
	fmt.Println(m)

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3
	fmt.Println(m)
	fmt.Println()

	fmt.Println(m["k1"], m["k2"])

	delete(m, "k2")
	fmt.Println(m["k1"], m["k2"])

	a, prs := m["k3"]
	fmt.Println("does k3 exists", prs, a)
	fmt.Println()

	n := map[string]int{"foo": 10, "bar": 2}
	fmt.Println(n)
	fmt.Println()

	msg, err := fmt.Println("1234567890")
	fmt.Println(msg, err)
}
