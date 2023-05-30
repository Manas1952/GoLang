package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"

	// it's necessary because in for loop it will wait for next element from channel which is not send, so error will be generated
	close(queue)

	// it internally receives the value (<- queue)
	for element := range queue {
		fmt.Println(element)
	}
}
