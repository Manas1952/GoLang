package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) person {
	p := person{name: name, age: 21}
	p.age = 22

	//fmt.Println(reflect.TypeOf(p))

	return p
}

func newPerson1(name string) *person {
	p := person{name: name, age: 21}
	p.age = 22

	//fmt.Println(reflect.TypeOf(p))

	return &p
}

func main() {
	fmt.Println(person{"Manas", 21})
	fmt.Println(person{name: "Manas", age: 21})
	fmt.Println(person{name: "Manas"})
	fmt.Println(&person{name: "Manas", age: 21})

	fmt.Println(newPerson("Purohit"))

	s1 := person{"p1", 12}
	fmt.Println(s1.name, s1.age)

	s2 := &s1
	fmt.Println(s2.age)

	s2.age = 13
	fmt.Println(s1.age)
}
