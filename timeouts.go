package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "result 1"
	}()

	select {
	case result := <-c1:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		c2 <- "result 2"
	}()

	select {
	case result := <-c2:
		fmt.Println(result)
	case x := <-c2:
		fmt.Println(x)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
