package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() {
	r.width = 11
	r.height = 22
	//return p := 2 * (r.width + r.height)
}

func (r rect) perimeter() {
	r.width = 33
	r.height = 44
	//return p := r.height * r.width
}

func main() {
	r := rect{1, 2}

	r.perimeter()
	fmt.Println(r)

	r.area()
	fmt.Println(r)

}
