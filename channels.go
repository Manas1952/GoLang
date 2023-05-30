package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messages <- "ping"
		messages <- "ping1"
	}()

	var msg string = <-messages
	var msg1 string = <-messages
	fmt.Println(msg)
	fmt.Println(msg1)
}
