package mutex

import (
	"fmt"
	"sync"

	"github.com/aafriya/go-playground/utils"
)

var balance = 0

func deposit(amount int, wg *sync.WaitGroup) {
	balance += amount
	wg.Done()
}

func Mutex() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go deposit(1, &wg) // deposit $1
	}

	wg.Wait()
	fmt.Println("Final Balance:", balance)
}

func Test() {
	utils.FromA()
	utils.FromB()
}

func init() {
	fmt.Println("Mutex package initialized")
}

func init() {
	fmt.Println("Mutex package initialized again")
}
