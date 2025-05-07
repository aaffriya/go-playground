package mutex

import (
	"fmt"
	"sync"
)

var Balance uint16

func depositWithoutMutex(amount *uint16, wg *sync.WaitGroup) {
	*amount += 1
	wg.Done()
}

func depositWithMutex(amount *uint16, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	*amount += 1
	mu.Unlock()
	wg.Done()
}

func Mutex() {
	var wg sync.WaitGroup

	// Without mutex
	for range 1000 {
		wg.Add(1)
		go depositWithoutMutex(&Balance, &wg)
	}
	wg.Wait() // ⏳ Wait for all goroutines
	fmt.Println("Final Balance without mutex is:", Balance)

	// Reset for next round
	Balance = 0
	var mu sync.Mutex
	wg = sync.WaitGroup{} // Reset WaitGroup

	// With mutex
	for range 1000 {
		wg.Add(1)
		go depositWithMutex(&Balance, &mu, &wg)
	}
	wg.Wait() // ⏳ Wait here too!
	fmt.Println("Final Balance with mutex is:", Balance)
}
