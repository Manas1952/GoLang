package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case signal := <-signals:
		fmt.Println("received signal", signal)
	default:
		fmt.Println("no activity")
	}
}
