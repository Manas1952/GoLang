package main

import "fmt"

func main() {
	messages := make(chan string, 4)

	messages <- "msg1"
	messages <- "msg2"
	messages <- "msg3"
	messages <- "msg3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
