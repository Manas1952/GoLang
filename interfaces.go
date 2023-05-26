package main

import "fmt"

type rectangle struct {
	width, height int
}

type circle struct {
	radius int
}

type shape interface {
	area() int
	perimeter() int
}

func (r rectangle) area() int {
	return r.height * r.width
}

//func (r rectangle) perimeter() int {
//	return 2 * (r.width + r.height)
//}

func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	//fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{1, 2}

	printData(r)
}
