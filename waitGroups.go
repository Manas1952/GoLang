package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var waitGroup sync.WaitGroup

	//i:= 10
	//_ = i

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)

		i := i

		go func() {
			defer waitGroup.Done()

			worker(i)
		}()
	}

	// prevents main function to end
	waitGroup.Wait()
}
