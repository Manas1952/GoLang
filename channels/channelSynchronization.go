package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	// will wait for go routine to end
	<-done
	fmt.Println("Main ended")
}
