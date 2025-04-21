package channel

import (
	"fmt"
	"time"
)

func ForLoop() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Simulate sending messages from different goroutines
	go func() {
		// time.Sleep(1 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from ch2"
	}()

	// close(ch1)
	// var i uint64 = 0

	// Loop with select to handle multiple channels
	for {
		select {
		case msg1, ok := <-ch1:
			fmt.Println("Received1:", msg1, " ok: ", ok)
		case msg2 := <-ch2:
			fmt.Println("Received2:", msg2)
			// return // Exit after receiving from ch2 (optional)
		// case <-time.After(3 * time.Second):
		// 	fmt.Println("Timeout!")
			// return
			// default:
			// 	if i%1000000 == 0 {
			// 		fmt.Println("No messages received, waiting...", i)
			// 	}
			// 	i++
		}
	}
}
