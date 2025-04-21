package main

import (
	. "fmt"
	// "time"
	"github.com/aaffriya/go-playground/utils"
	// "github.com/aaffriya/go-playground/utils/myutils"
	// "github.com/aaffriya/go-playground/utils/myutils/error"
	// "github.com/aaffriya/go-playground/channel"
	// _defer "github.com/aaffriya/go-playground/defer"
	mutex "github.com/aaffriya/go-playground/mutex"
	// _ "github.com/aaffriya/go-playground/mutex"
	"github.com/google/uuid"
)

func main() {
	Println("Program started")

	id := uuid.New()
	Println("Generated UUID:", id)
	// chann := make(chan int, 1)
	// fmt.Println(<-chann)

	// result := utils.Add(5, 3)
	// fmt.Println("Result:", result)

	// utils.Subtract(5, 3)
	// fmt.Println("Result:", result)

	// myutils.PrintHelloFromMyUtils()

	// errorHandling.Program()

	// fmt.Println("calling channel.Main 1")
	// go channel.Channel(1)
	// fmt.Println("calling channel.Main 2")
	// go channel.Main(2)
	// fmt.Println("calling channel.Main 3")
	// go channel.Main(3)
	// fmt.Println("calling channel.Main 4")
	// go channel.Main(4)
	// fmt.Println("calling channel.Main 5")
	// go channel.Main(5)
	// fmt.Println("calling channel.Main 6")
	// go channel.Main(6)
	// fmt.Println("calling channel.Main 7")
	// go channel.Main(7)
	// fmt.Println("calling channel.Main 8")
	// go channel.Main(8)
	// fmt.Println("calling channel.Main 9")
	// go channel.Main(9)

	// channel.Select()

	// channel.ForLoop()
	// channel.Reflect()

	// _defer.Example()

	// time.Sleep(time.Second * 2)

	// _defer.TestDefer()
	mutex.Mutex()

	utils.FromA()
	utils.FromB()
}
