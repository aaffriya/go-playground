package channel

import (
	"fmt"
	"time"
)

func Select() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	// Start goroutine 2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// Select waits for the first available channel
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	// case <-time.After(2 * time.Second):
	// 	fmt.Println("Timeout")
	default:
		fmt.Println("No messages received")
	}
}
