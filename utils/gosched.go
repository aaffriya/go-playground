package utils

import (
	"fmt"
	"sync"
	"runtime"
)

func busyWait(str string, wg *sync.WaitGroup) {
	// Simulate a busy wait
	for i := range 10 {
		fmt.Println(str, " ", i)
		runtime.Gosched() // Yield the processor
	}
	wg.Done() // Signal that this goroutine is done
}

func Suspend() {
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(4) // Add 4 goroutines to the WaitGroup
	go busyWait("Goroutine 1", &wg)
	go busyWait("Goroutine 2", &wg)
	go busyWait("Goroutine 3", &wg)
	go busyWait("Goroutine 4", &wg)
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines have finished.")
}