package _struct

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func StructExample() {
	p := Person{Name: "John", Age: 30}
	p.Greet()
}