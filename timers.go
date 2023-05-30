package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop := timer2.Stop()

	if stop {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
