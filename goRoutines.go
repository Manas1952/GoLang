package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(s, ":", i)
	}
}

func main() {
	f("direct")

	go f("go routine")
	go f("go routine 1")

	go func(str string) {
		fmt.Println(str)
	}("anonymous go routine")

	time.Sleep(time.Second * 3)
	fmt.Println("done")
}
