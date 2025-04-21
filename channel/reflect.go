package channel

import (
	"fmt"
	"reflect"
	"time"
	"slices"
)

func Reflect() {
	// Create 3 channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	// Start goroutines that send messages with different delays
	go func() { time.Sleep(1 * time.Second); ch1 <- "from ch1"; close(ch1) }()
	go func() { time.Sleep(2 * time.Second); ch2 <- "from ch2"; close(ch2) }()
	go func() { time.Sleep(3 * time.Second); ch3 <- "from ch3"; close(ch3) }()

	// Create a slice of reflect.SelectCase to represent select cases
	cases := []reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,   // Direction: receive
			Chan: reflect.ValueOf(ch1), // Channel 1
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch2), // Channel 2
		},
		{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch3), // Channel 3
		},
	}

	// Listen until all channels are closed
	for len(cases) > 0 {
		// Wait for any channel to receive
		chosen, recv, ok := reflect.Select(cases)

		if !ok {
			// Channel was closed, remove it from the list
			fmt.Println("Channel", chosen, "was closed")
			cases = slices.Delete(cases, chosen, chosen+1)
			continue
		}

		// Print the received value and which channel it came from
		fmt.Println("Received:", recv.String(), "from channel", chosen)
	}
}
