package _defer

import (
	"fmt"
	"testing"
	"time"
)

func Defer() string {
	// Defer is used to ensure that a function call is performed later in a program's execution
	// Defer statements are executed in LIFO order (Last In, First Out)
	// Defer is often used for cleanup tasks, such as closing files or releasing resources

	// Example 1: Basic usage of defer
	defer fmt.Println("World")
	fmt.Println("Hello")

	// Example 2: Defer with multiple statements
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

	// Example 3: Defer with a function call
	defer func() {
		fmt.Println("Deferred function executed")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function executed")
	// Example 4: Defer with a return statement
	defer fmt.Println("This will be printed before the return statement")
	return "Defer example completed"
}

func Example() {

	x := "Hello"
	y := &x
	defer fmt.Println("Deferred:", *y)
	x = "World"
	x = "Hello, World"
	fmt.Println("Changed:", x)
}

func doNothing() {}

func BenchmarkDirectCall(b *testing.B) {
	for b.Loop() {
		doNothing()
	}
}

func BenchmarkDeferCall(b *testing.B) {
	for b.Loop() {
		defer doNothing()
	}
}

func TestDefer() {
	benchmarkDirect := testing.Benchmark(BenchmarkDirectCall)
	benchmarkDiffer := testing.Benchmark(BenchmarkDeferCall)
	fmt.Println("BenchmarkDirectCall:", benchmarkDirect)
	fmt.Println("BenchmarkDeferCall:", benchmarkDiffer)
}
