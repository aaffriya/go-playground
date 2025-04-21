package channel

import (
	"fmt"
	// "time"
)

func fanIn(ch1, ch2 <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        for {
            select {
            case msg := <-ch1:
                out <- msg
            case msg := <-ch2:
                out <- msg
            }
        }
    }()
    return out
}

func Combine() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        for i := range 3 {
            ch1 <- fmt.Sprintf("from ch1: %d", i)
            // time.Sleep(100 * time.Millisecond)
        }
    }()
    
    go func() {
        for i := range 3 {
            ch2 <- fmt.Sprintf("from ch2: %d", i)
            // time.Sleep(150 * time.Millisecond)
        }
    }()

    out := fanIn(ch1, ch2)

    for range 6 {
        fmt.Println(<-out)
    }
}
