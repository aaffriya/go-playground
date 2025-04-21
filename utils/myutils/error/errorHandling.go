package errorHandling

import (
    // "errors"
    "fmt"
)

// func divide(a, b float64) (float64, error) {
//     if b == 0 {
//         return 0, errors.New("cannot divide by zero")
//     }
//     return a / b, nil
// }

func Program() {
    err := doSomething()
    fmt.Println("Error:", err)
}

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 500, Message: "Something went wrong"}
}

