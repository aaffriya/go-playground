package channel

import (
	"fmt"
	// "time"
)

func Bidirectional() {
	ch := make(chan string)

	// Start a goroutine to send messages
	go func() {
		for i := range 5 {
			ch <- fmt.Sprintf("Message %d", i)
			// time.Sleep(1 * time.Second)
		}
		close(ch) // Close the channel when done
	}()

	// Receive messages from the channel
	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}