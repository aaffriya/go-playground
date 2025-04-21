package channel

import (
	"fmt"
	// "time"
)

func Channel(number int) {
	// chann := make(chan int, 1)
	// fmt.Println(<-chann)
	ch := make(chan string)

	fmt.Println("Starting goroutine " + fmt.Sprintf("%d", number))
	go func(number int, ch chan string) {
		// time.Sleep(time.Second * 2)
		fmt.Println("Inside goroutine " + fmt.Sprintf("%d", number))
		ch <- "Hello from the goroutine! " + fmt.Sprintf("%d", number)
		fmt.Println("Inside goroutine " + fmt.Sprintf("%d", number) + " after sending first message")
		// time.Sleep(time.Second * 2)
		ch <- "Second message from the goroutine! " + fmt.Sprintf("%d", number)
		fmt.Println("Inside goroutine " + fmt.Sprintf("%d", number) + " after sending second message")
		ch <- "Second message from the goroutine! " + fmt.Sprintf("%d", number)
		fmt.Print("all done\n")
	}(number, ch)

	fmt.Println("Goroutine " + fmt.Sprintf("%d", number) + " started")
	message := <-ch

	fmt.Println(message)
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
// 	fmt.Println("Goroutine " + fmt.Sprintf("%d", number) + " received first message")
// 	message = <-ch

// 	fmt.Println(message)
// 	fmt.Println("Goroutine " + fmt.Sprintf("%d", number) + " received first message")
// 	message = <-ch

// 	fmt.Println(message)
	fmt.Println("Goroutine " + fmt.Sprintf("%d", number) + " finished")
}
